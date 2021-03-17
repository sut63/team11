import React, { useState } from 'react';
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
  position: 'center',
  showConfirmButton: false,
  showCloseButton: true,

});

export default function ComponentsTable() {

  const classes = useStyles();
  const api = new DefaultApi();
  const [loading, setLoading] = useState(true);
  const [search, setSearch] = useState(true);

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
    setBook([]);
  }

  const SearchBook = async () => {
    if (bookname == "") {
      alertMessage("info", "แสดงข้อมูลหนังสือทั้งหมดในระบบ")
      const apiUrl = `http://localhost:8080/api/v1/searchbooks?book=${bookname}`;
      const requestOptions = {
        method: 'GET',
      };
      fetch(apiUrl, requestOptions)
        .then(response => response.json())
        .then(data => {
          console.log(data.data)
          if (data.data != null) {
            if (data.data.length >= 1) {
              console.log(data.data)
              setBook(data.data);
            }
          }
        });

    }
    else {
      const apiUrl = `http://localhost:8080/api/v1/searchbooks?book=${bookname}`;
      const requestOptions = {
        method: 'GET',
      };
      fetch(apiUrl, requestOptions)
        .then(response => response.json())
        .then(data => {
          console.log(data.data)
          alertMessage("warning", "ไม่พบข้อมูลหนังสือที่ค้นหา")
          setBook([]);
          if (data.data != null) {
            if (data.data.length >= 1) {
              alertMessage("success", "พบข้อมูลหนังสือที่ค้นหา")
              console.log(data.data)
              setBook(data.data);
            }
          }
        });

    }
  }


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
                    SearchBook();

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
                      <TableCell align="center">สถานะหนังสือ</TableCell>
                    </TableRow>
                  </TableHead>
                  <TableBody>

                    {book.map((item: any) => (
                      <TableRow key={item.id}>
                        <TableCell align="center">{item.id}</TableCell>
                        <TableCell align="center">{item.BookName}</TableCell>
                        <TableCell align="center">{item.BookPage}</TableCell>
                        <TableCell align="center">{item.Barcode}</TableCell>
                        <TableCell align="center">{item.edges?.Author?.Name}</TableCell>
                        <TableCell align="center">{item.edges?.Category?.CategoryName}</TableCell>
                        <TableCell align="center">{item.edges?.Status?.STATUS_NAME}</TableCell>
                      </TableRow>
                    ))}
                  </TableBody>
                </Table>
              </TableContainer>

            </Paper>
          </Grid>
        </Grid>
      </Content>
    </Page>
  );

}
