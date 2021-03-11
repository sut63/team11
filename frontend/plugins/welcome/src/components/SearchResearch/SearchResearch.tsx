import React, { useState} from 'react';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';

import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import { Link as RouterLink } from 'react-router-dom';

import { EntResearch } from '../../api/models/EntResearch';

import moment from 'moment';
import { Page, pageTheme, Header, Content, Link } from '@backstage/core';
import { Grid, Button, TextField, Typography, FormControl } from '@material-ui/core';
import SearchTwoToneIcon from '@material-ui/icons/SearchTwoTone';
import { Alert } from '@material-ui/lab';

import { styled } from '@material-ui/core/styles';
import { compose, spacing, palette, sizing, shadows   } from '@material-ui/system';

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

  const profile = { givenName: 'ยินดีต้อนรับสู่ ระบบค้นหาข้อมูลงานวิจัย' };
  const name = JSON.parse(String(localStorage.getItem("userName")));
  const userName = name;
  const classes = useStyles();

  const [research, setResearch] = useState<EntResearch[]>([])
  const [alert, setAlert] = useState(true);
  const [status, setStatus] = useState(false);
  const [docname, setdocname] = useState(String);

  const Box = styled('div')(compose(spacing, palette, shadows, sizing ));
  
  //-------------------
  const Docnamehandlehange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setStatus(false);
    setdocname(event.target.value as string);
  };

  const cleardata = () => {
    setdocname("");
    setStatus(false)
    setResearch([]);
  }

  //---------------------
  const SearchResearch = async () => {
      setStatus(true);
      setAlert(true);
      const apiUrl = `http://localhost:8080/api/v1/searchresearchs?research=${docname}`;
      const requestOptions = {
        method: 'GET',
      };
      fetch(apiUrl, requestOptions)
        .then(response => response.json())
        .then(data => {
          console.log(data.data)
          setStatus(true);
          setAlert(false);
          setResearch([]);
          if (data.data != null) {
            if (data.data.length >= 1) {
              setStatus(true);
              setAlert(true);
              console.log(data.data)
              setResearch(data.data);
            }
          }
        });
  }

  return (
    <Page theme={pageTheme.home}>
      <Header title={`${profile.givenName}`}>
        <table>
          <tr>
            <th>
              <Box color="black" bgcolor="info.main" p={1} boxShadow={1} width="auto">
                <font size='5' color= "black">Hello are you 🌵 {userName}🌵</font>
              </Box>
            </th>
          </tr>
          <tr>
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
              <Typography align="center">
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
                  <FormControl className={classes.margin} variant="outlined">
                    <div className={classes.paper}><font size='5'>ชื่องานวิจัย 📜 </font></div>
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
                  onClick={() => {SearchResearch();}}

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
                  onClick={() => {cleardata();}}
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

          </Grid>
        </Grid>


        <Grid container justify="center">
          <Grid item xs={12} md={10}>
            <Paper>

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
                        <TableCell align="center">{item.DOC_NAME}</TableCell>
                        <TableCell align="center">{item.PAGE_NUMBER}</TableCell>
                        <TableCell align="center">{item.YEAR_NUMBER}</TableCell>
                        <TableCell align="center">{item.edges?.MyDoc?.Name}</TableCell>
                        <TableCell align="center">{item.edges?.DocType?.TYPE_NAME}</TableCell>
                        <TableCell align="center">{moment(item.DATE).format('DD/MM/YYYY HH.mm น.')}</TableCell>
                      </TableRow>
                    ))}
                  </TableBody>
                </Table>
              </TableContainer>

            </Paper>
          </Grid>
        </Grid>
      </Content >
    </Page >
  );

}
