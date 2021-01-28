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

import { EntBook } from '../../api/models/EntBook';

import Swal from 'sweetalert2'
import { Link as RouterLink } from 'react-router-dom';
import { Page, pageTheme, Header, Content, Link } from '@backstage/core';
import { Grid, Button, TextField, Typography, FormControl } from '@material-ui/core';
import SearchTwoToneIcon from '@material-ui/icons/SearchTwoTone';


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
const Toast = Swal.mixin({
  // toast: true,
  position: 'center',
  showConfirmButton: false,
  //timer: 3000,
  //timerProgressBar: true,
  showCloseButton: true,

});

export default function ComponentsTable() {

    //const name = JSON.parse(String(localStorage.getItem("userName")));
   // const userName = name
      //--------------------------
      ;
  
    const classes = useStyles();
    const api = new DefaultApi();
    const [loading, setLoading] = useState(true);
    const [search, setSearch] = useState(false);
  
    //---------------------------
    const [checkbookname, setbooknames] = useState(false);
    const [book, setBook] = useState<EntBook[]>([])
  
    //--------------------------
    const [bookname, setbookname] = useState(String);
    const profile = { givenName: 'ยินดีต้อนรับสู่ ระบบค้นหาข้อมูลหนังสือ' };
    const alertMessage = (icon: any, title: any) => {
      Toast.fire({
        icon: icon,
        title: title,
      });
      setSearch(false);
    }
  
    useEffect(() => {
      const getBooks = async () => {
        const res = await api.listBook({ offset: 0 });
        setLoading(false);
        setBook(res);
      };
      getBooks();
    }, [loading]);
  
    //-------------------
    const Booknamehandlehange = (event: React.ChangeEvent<{ value: unknown }>) => {
      setSearch(false);
      setbooknames(false);
      setbookname(event.target.value as string);
  
    };
  
    const cleardata = () => {
      setbookname("");
      setSearch(false);
      setbooknames(false);
      setSearch(false);
  
    }
    //---------------------
    const checkresearch = async () => {
      var check = false;
      book.map(item => {
        if (bookname != "") {
          if (item.bookName?.includes(bookname)) {
            setbooknames(true);
            alertMessage("success", "พบข้อมูลหนังสือที่ค้นหา");
            check = true;
          }
        }
      })
      if (!check) {
        alertMessage("error", "ไม่พบข้อมูลหนังสือที่ค้นหา");
      }
      console.log(checkbookname)
      if (bookname == "") {
        alertMessage("info", "แสดงข้อมูลหนังสือทั้งหมดในระบบ");
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
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
            <Link component={RouterLink} to="/">
            <Button variant="contained" style={{ background: '#FFFFFF', height: 36 }}>
                  <h3
                    style={
                      {
                        color: "#000000",
                        borderRadius: 10,
                        height: 25,
                        padding: '0 5px',
                      }
                    }>
                    ย้อนกลับ
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
                  <div style={{ background: '#BC85A3', height: 55 }}>
                    <h1 style={
                      {
                        color: "#FFFFFF",
                        borderRadius: 5,
                        height: 18,
                        padding: '0 5px',
                        fontSize: '36px',
                      }}>
                      ค้นหาข้อมูลหนังสือ
              </h1>
                  </div>
  
                  <div>
                    <FormControl
                      className={classes.margin}
                      variant="outlined"
                    >
                      <div className={classes.paper}><strong>กรุณากรอกชื่อหนังสือ </strong></div>
                      <TextField
                        id="bookname"
                        value={bookname}
                        onChange={Booknamehandlehange}
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
                      ค้นหา
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
                      ลบการค้นหาทั้งหมด
              </h3>
                  </Button>
                </Typography>
              </Paper>
            </Grid>
          </Grid>
  
  
          <Grid container justify="center">
            <Grid item xs={12} md={10}>
              <Paper>
                {search ? (
                  <div>
                    {  checkbookname ? (
                      <TableContainer component={Paper}>
                        <Table className={classes.table} aria-label="simple table">
                          <TableHead>
                            <TableRow>
                              <TableCell align="center">ลำดับที่</TableCell>
                              <TableCell align="center">ชื่อหนังสือ</TableCell>
                              <TableCell align="center">จำนวนหน้า</TableCell>
                              <TableCell align="center">รหัสBarcode</TableCell>
                              <TableCell align="center">ชื่อผู้แต่ง</TableCell>
                              <TableCell align="center">หมวด</TableCell>
                            </TableRow>
                          </TableHead>
                          <TableBody>
  
                            {book.filter((filter: any) => filter.bookName.includes(bookname)).map((item: any) => (
                              <TableRow key={item.id}>
                                <TableCell align="center">{item.id}</TableCell>
                                <TableCell align="center">{item.bookName}</TableCell>
                                <TableCell align="center">{item.bookPage}</TableCell>
                                <TableCell align="center">{item.barcode}</TableCell>
                                <TableCell align="center">{item.edges?.author?.name}</TableCell>
                                <TableCell align="center">{item.edges?.category?.categoryName}</TableCell>
                              </TableRow>
                            ))}
                          </TableBody>
                        </Table>
                      </TableContainer>
                    )
                      : bookname == "" ? (
                        <div>
                          <TableContainer component={Paper}>
                            <Table className={classes.table} aria-label="simple table">
                              <TableHead>
                                <TableRow>
                                  <TableCell align="center">ลำดับที่</TableCell>
                                  <TableCell align="center">ชื่อหนังสือ</TableCell>
                                  <TableCell align="center">จำนวนหน้า</TableCell>
                                  <TableCell align="center">รหัสBarcode</TableCell>
                                  <TableCell align="center">ชื่อผู้แต่ง</TableCell>
                                  <TableCell align="center">หมวด</TableCell>
                                </TableRow>
                              </TableHead>
                              <TableBody>
  
                                {book.map((item: any) => (
                                  <TableRow key={item.id}>
                                    <TableCell align="center">{item.id}</TableCell>
                                    <TableCell align="center">{item.bookName}</TableCell>
                                    <TableCell align="center">{item.bookPage}</TableCell>
                                    <TableCell align="center">{item.barcode}</TableCell>
                                    <TableCell align="center">{item.edges?.author?.name}</TableCell>
                                    <TableCell align="center">{item.edges?.category?.categoryName}</TableCell>
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
  