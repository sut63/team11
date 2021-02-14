import React, { useState, useEffect } from 'react';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import { DefaultApi } from '../../api/apis';

import { EntResearch } from '../../api/models/EntResearch';

import { Link as RouterLink } from 'react-router-dom';
import moment from 'moment';
import { Page, pageTheme, Header, Content, Link } from '@backstage/core';
import { Grid, Button, TextField, Typography, FormControl } from '@material-ui/core';
import SearchTwoToneIcon from '@material-ui/icons/SearchTwoTone';

import { Alert } from '@material-ui/lab';


const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    headsearch: {
      width: 'auto',
      margin: '10px',
      color: '#FFFFFF',
      background: '#2196F3',
    },
    margin: {
      margin: theme.spacing(1),
    },
    margins: {
      margin: theme.spacing(2),
    },

    withoutLabel: {
      marginTop: theme.spacing(3),
    },
    textField: {
      width: '25ch',
    },
    paper: {
      marginTop: theme.spacing(3),
      marginBottom: theme.spacing(3),
    },
    table: {
      minWidth: 500,
    },


  }),
);

export default function ComponentsTable() {

  const name = JSON.parse(String(localStorage.getItem("userName")));
  const userName = name
    //--------------------------
    ;

  const classes = useStyles();
  const api = new DefaultApi();
  const [loading, setLoading] = useState(true);
  const [search, setSearch] = useState(false);

  //---------------------------
  const [checkdocname, setdocnames] = useState(false);
  const [research, setResearch] = useState<EntResearch[]>([])

  const [alert, setAlert] = useState(true);
  const [status, setStatus] = useState(false);

  //--------------------------
  const [docname, setdocname] = useState(String);
  const profile = { givenName: 'ยินดีต้อนรับสู่ ระบบค้นหาข้อมูลงานวิจัย' };
  
  useEffect(() => {
    const getResearchs = async () => {
      const res = await api.listResearch({ offset: 0 });
      setLoading(false);
      setResearch(res);
    };
    getResearchs();
  }, [loading]);

  //-------------------
  const Docnamehandlehange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setSearch(false);
    setdocnames(false);
    setStatus(false);
    setdocname(event.target.value as string);

  };

  const cleardata = () => {
    setdocname("");
    setSearch(false);
    setdocnames(false);
    setSearch(false);
    setStatus(false)

  }
  //---------------------
  const checkresearch = async () => {
    var check = false;
    research.map(item => {
      if (docname != "") {
        if (item.dOCNAME?.includes(docname)) {
          setdocnames(true);
          setStatus(true);
          setAlert(true);

          //พบข้อมูลที่ค้นหา
          
          check = true;
        }
      }
    })
    if (!check) {
      setStatus(true);
      setAlert(false);

      //ไม่พบข้อมูลที่ค้นหา
    }
    console.log(checkdocname)
    if (docname == "" ) {
      setStatus(true);
      setAlert(true);

      //แสดงข้อมูลงานวิจัยทั้งหมดในระบบ
    }    
  };

  return (
    <Page theme={pageTheme.home}>
      <Header
        title={`${profile.givenName}`}
      >
        <table>
          <tr>
            <th>
              <Button variant="contained" color='primary' size='large' >
                <font size='3'>hello are you 🌵 {userName}🌵</font>
              </Button>
            </th>
            <th>
              
          <Link component={RouterLink} to="/">
                <Button variant="contained" style={{ background: 'linear-gradient(45deg, #3399FF 15%, #9900FF 120%)', height: 36 }}>
                  <h3
                    style={
                      {
                        color: "#FFFAF0",
                        borderRadius: 10,
                        height: 25,
                        padding: '0 20px',
                      }
                    }>
                    กลับหน้าหลัก
            </h3>
                </Button>
              </Link>
            </th>
          </tr>
        </table>

      </Header>
      <Content>
        <Grid container item xs={12} justify="center">
          <Grid item xs={5}>
            <Paper>

              <Typography align="center" >
                <div style={{ background: 'linear-gradient(45deg, #00FFCC 15%, #9900FF 120%)', height: 55 }}>
                  <h1 style={
                    {
                      color: "#000000",
                      borderRadius: 5,
                      height: 18,
                      padding: '0 30px',
                      fontSize: '40px',
                    }}>
                    ค้นหาข้อมูลงานวิจัย
            </h1>
                </div>

                <div>
                  <FormControl
                    className={classes.margin}
                    variant="outlined"
                  >
                    <div className={classes.paper}>
                    <font size='5'>ชื่องานวิจัย 📜 </font>
                     </div>
                    <TextField
                      id="docname"
                      value={docname}
                      onChange={Docnamehandlehange}
                      type="string"
                      size="small"

                      style={{ width: 200 }}
                    />
                  </FormControl>
                </div>
                <div></div>
                <Button
                  onClick={() => {
                    checkresearch();
                    setSearch(true);

                  }}
                  endIcon={<SearchTwoToneIcon />}
                  className={classes.margins}
                  variant="contained"
                  style={{ background: "#9900FF", height: 40 }}>
                  <h3
                    style={
                      {
                        color: "#FFFFFF",
                        padding: '0 10px',

                      }
                    }>
                    ค้นหาข้อมูล
            </h3>
                </Button>
                <Button
                  onClick={() => {
                    cleardata();

                  }}
                  className={classes.margins}
                  variant="contained"
                  style={{ background: "#9900FF", height: 40 }}>
                  <h3
                    style={
                      {
                        color: "#FFFFFF",
                        padding: '0 25px',

                      }
                    }>
                    เคลียร์ข้อมูล
            </h3>
                </Button>
              </Typography>
            </Paper>
          </Grid>
        </Grid>
        {status ? (
            <div>
              {alert ? (
                <Alert severity="success">
                  แสดงข้อมูลงานวิจัย {docname} 📜
                </Alert>
              )
              : (
                  <Alert severity="warning" style={{ marginTop: 20 }}>
                    ไม่พบข้อมูลที่ค้นหา
                  </Alert>
                )}
            </div>
          ) : null}
        
        <Grid container justify="center">
          <Grid item xs={12} md={10}>
            <Paper>
              
              {search ? (
                <div>
                  {  checkdocname ? (
                    <TableContainer component={Paper}>
                      <Table className={classes.table} aria-label="simple table">
                        <TableHead>
                          <TableRow>
                            <TableCell align="center">ลำดับที่</TableCell>
                            <TableCell align="center">ชื่องานวิจัย</TableCell>
                            <TableCell align="center">จำนวนหน้า</TableCell>
                            <TableCell align="center">ปีที่พิมพ์</TableCell>
                            <TableCell align="center">ชื่อผู้แต่ง</TableCell>
                            <TableCell align="center">ประเภทงานวิจัย</TableCell>
                            <TableCell align="center">วันที่</TableCell>
                          </TableRow>
                        </TableHead>
                        <TableBody>

                          {research.filter((filter: any) => filter.dOCNAME.includes(docname)).map((item: any) => (
                            <TableRow key={item.id}>
                              <TableCell align="center">{item.id}</TableCell>
                              <TableCell align="center">{item.dOCNAME}</TableCell>
                              <TableCell align="center">{item.pAGENUMBER}</TableCell>
                              <TableCell align="center">{item.yEARNUMBER}</TableCell>
                              <TableCell align="center">{item.edges?.myDoc?.name}</TableCell>
                              <TableCell align="center">{item.edges?.docType?.tYPENAME}</TableCell>
                              <TableCell align="center">{moment(item.dATE).format('DD/MM/YYYY HH.mm น.')}</TableCell>
                            </TableRow>
                          ))}
                        </TableBody>
                      </Table>
                    </TableContainer>
                  )
                  
                    : docname == "" ? (
                      
                      <div>
                        <TableContainer component={Paper}>
                          <Table className={classes.table} aria-label="simple table">
                            <TableHead>
                              <TableRow>
                                <TableCell align="center">ลำดับที่</TableCell>
                                <TableCell align="center">ชื่องานวิจัย</TableCell>
                                <TableCell align="center">จำนวนหน้า</TableCell>
                                <TableCell align="center">ปีที่พิมพ์</TableCell>
                                <TableCell align="center">ชื่อผู้แต่ง</TableCell>
                                <TableCell align="center">ประเภทงานวิจัย</TableCell>
                                <TableCell align="center">วันที่</TableCell>
                              </TableRow>
                            </TableHead>
                            <TableBody>

                              {research.map((item: any) => (
                                <TableRow key={item.id}>
                                  <TableCell align="center">{item.id}</TableCell>
                                  <TableCell align="center">{item.dOCNAME}</TableCell>
                                  <TableCell align="center">{item.pAGENUMBER}</TableCell>
                                  <TableCell align="center">{item.yEARNUMBER}</TableCell>
                                  <TableCell align="center">{item.edges?.myDoc?.name}</TableCell>
                                  <TableCell align="center">{item.edges?.docType?.tYPENAME}</TableCell>
                                  <TableCell align="center">{moment(item.dATE).format('DD/MM/YYYY HH.mm น.')}</TableCell>
                                </TableRow>
                              ))}
                            </TableBody>
                          </Table>
                        </TableContainer>

                      </div>
                    ) : null}
                </div>
              ) : null}
            </Paper>
          </Grid>
        </Grid>
      </Content>
    </Page>
  );

}
