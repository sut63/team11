import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import Swal from 'sweetalert2';
import { DefaultApi } from '../../api/apis';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';
import { EntUser } from '../../api/models/EntUser';
import { EntBook } from '../../api/models/EntBook'; 
import { EntServicePoint } from '../../api/models/EntServicePoint'; 


const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    margin: {
      margin: theme.spacing(3),
    },

    withoutLabel: {
      marginTop: theme.spacing(3),
    },
    textField: {
      width: '25ch',
    },
    paper: {
      marginTop: theme.spacing(1),
      marginBottom: theme.spacing(1),
      marginLeft: theme.spacing(1),
    },
    formControl: {
      width: 400,
    },
  }),
);

const Toast = Swal.mixin({
  toast: true,
  position: 'top-end',
  width: '400px',
  padding: '100px',
  showConfirmButton: false,
  timer: 3000,
  timerProgressBar: true,
  didOpen: toast => {
    toast.addEventListener('mouseenter', Swal.stopTimer);
    toast.addEventListener('mouseleave', Swal.resumeTimer);
  },
});

export default function Create() {
  const classes = useStyles();
  const api = new DefaultApi();

  const [users, setUsers] = useState<EntUser[]>([]);
  //-------
  const [books, setBooks] = useState<EntBook[]>([]);
  const [servicepoints, setServicepoints] = useState<EntServicePoint[]>([]);

  //-------
  const [userID, setUser] = useState(Number);
  const [bookID, setBook] = useState(Number);
  const [servicePointID, setServicepoint] = useState(Number);


  const [loading, setLoading] = useState(true);


  useEffect(() => {
    //-------
    const getBook = async () => {
      const res = await api.listBook({ limit: 10, offset: 0 });
      setLoading(false);
      setBooks(res);
      console.log(res);
    };
    getBook();


    const getServicepoint = async () => {
      const res = await api.listServicepoint({ limit: 10, offset: 0 });
      setLoading(false);
      setServicepoints(res);
      console.log(res);
    };
    getServicepoint();

    const getUser = async () => {
      const res = await api.listUser({ limit: 10, offset: 0 });
      setLoading(false);
      setUsers(res);
      console.log(res);
    };
    getUser();

  }, [loading]);

  //----------------
  const handleBookchange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setBook(event.target.value as number);
  };

  const handleServicepointchange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setServicepoint(event.target.value as number);
  };

  const handleUserchange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setUser(event.target.value as number);
  };

      const bookBorrow = {
        bookID: bookID,
        servicePointID: servicePointID,
        userID: userID,
      };
      console.log(bookBorrow);
      function save() {
        const apiUrl = 'http://localhost:8080/api/v1/bookborrows';
        const requestOptions = {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify(bookBorrow),
        };

        console.log(bookBorrow); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

        fetch(apiUrl, requestOptions)
          .then(response => response.json())
          .then(data => {
            console.log(data);
            if (data.id != null) {
              //clear();
              Toast.fire({
                icon: 'success',
                title: 'บันทึกข้อมูลสำเร็จ',
              });
            } else {
              Toast.fire({
                icon: 'error',
                title: '<br><h5>กรุณากรอกข้อมูลใหม่</h5></br>',
              });
            }
          });
      }
  return (
    <Page theme={pageTheme.home}>
      <Header
        title={`ระบบยืมหนังสือ`}
        subtitle="กรุณากรอกข้อมูล"
      ></Header>

      <Content>
        <ContentHeader title="เพิ่มข้อมูลการยืมหนังสือ">

        </ContentHeader>

        <div className={classes.root}>
          <form noValidate autoComplete="off">

            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <div className={classes.paper}><strong>หนังสือ</strong></div>
                <InputLabel id="author-label"></InputLabel>
                <Select
                  labelId="ชื่อหนังสือ"
                  id="bookID"
                  value={bookID}
                  onChange={handleBookchange}
                  style={{ width: 400 }}
                >
                  {books.map((item: EntBook) => (
                    <MenuItem value={item.id}>{item.bookName}</MenuItem>
                  ))}
                </Select>
              </FormControl>
            </div>

            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <div className={classes.paper}><strong>ผู้ยืม</strong></div>
                <InputLabel id="author-label"></InputLabel>
                <Select
                  labelId="ผู้ยืม"
                  id="userID"
                  value={userID}
                  onChange={handleUserchange}
                  style={{ width: 400 }}
                >
                  {users.map((item: EntUser) => (
                    <MenuItem value={item.id}>{item.uSERNAME}</MenuItem>
                  ))}
                </Select>
              </FormControl>
            </div>

            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <div className={classes.paper}><strong>จุดบริการ</strong></div>
                <InputLabel id="researchtype-label"></InputLabel>
                <Select
                  labelId="จุดบริการที่ยืมหนังสือ"
                  id="servicepointID"
                  value={servicePointID}
                  onChange={handleServicepointchange}
                  style={{ width: 400 }}
                >
                  {servicepoints.map((item: EntServicePoint) => (
                    <MenuItem value={item.id}>{item.cOUNTERNUMBER}</MenuItem>
                  ))}
                </Select>
              </FormControl>
            </div>

            <div className={classes.margin}>
              <Button
                onClick={() => {
                  save();
                }}
                variant="contained"
                color="primary"
              >
                ยืนยันการบันทึก
             </Button>
              <Button
                style={{ marginLeft: 20 }}
                component={RouterLink}
                to="/"
                variant="contained"
              >
                ย้อนกลับ
             </Button>
            </div>
          </form>
        </div>
      </Content>
    </Page>
  );
}