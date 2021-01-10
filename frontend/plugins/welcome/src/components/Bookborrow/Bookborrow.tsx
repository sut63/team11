import React, { FC, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import Swal from 'sweetalert2'; // alert

import {
  Container,
  Grid,
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  TextField,
  Avatar,
  Button,
} from '@material-ui/core';
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
import {
    EntBook,
    EntServicePoint,
    EntUser,
  } from '../../api/models/';

// header css
const HeaderCustom = {
  minHeight: '50px',
};

// css style
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
  },
  formControl: {
    width: 300,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  textField: {
    width: 300,
  },
}));

interface Bookborrow {
  book: number;
  user: number;
  servicepoint: number;
  added: Date;
  // create_by: number;
}

const Bookborrow: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();

  const [bookborrow, setBookborrow] = React.useState<
    Partial<Bookborrow>
  >({});

  const [users, setUsers] = React.useState<EntUser[]>([]);
  const [books, setBooks] = React.useState<EntBook[]>([]);
  const [servicepoints, setServicepoints] = React.useState<EntServicePoint[]>([]);
  

  // alert setting
  const Toast = Swal.mixin({
    toast: true,
    position: 'top-end',
    showConfirmButton: false,
    timer: 3000,
    timerProgressBar: true,
    didOpen: toast => {
      toast.addEventListener('mouseenter', Swal.stopTimer);
      toast.addEventListener('mouseleave', Swal.resumeTimer);
    },
  });

  const getUsers = async () => {
    const res = await http.listUser({ limit: 10, offset: 0 });
    setUsers(res);
  };

  const getBook = async () => {
    const res = await http.listBook({ limit: 10, offset: 0 });
    setBooks(res);
  };

  const getServicepoint = async () => {
    const res = await http.listServicepoint({ limit: 10, offset: 0 });
    setServicepoints(res);
  };



  // Lifecycle Hooks
  useEffect(() => {
    getUsers();
    getBook();
    getServicepoint();
  }, []);

  // set data to object playlist_video
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,
  ) => {
    const name = event.target.name as keyof typeof Bookborrow;
    const { value } = event.target;
    setBookborrow({ ...bookborrow, [name]: value });
    console.log(bookborrow);
  };

  // clear input form
  function clear() {
    setBookborrow({});
  }

  // function save data
  function save() {
    const apiUrl = 'http://localhost:8080/api/v1/bookborrow';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(bookborrow),
    };

    console.log(bookborrow); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.status === true) {
          clear();
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        } else {
          Toast.fire({
            icon: 'error',
            title: 'บันทึกข้อมูลไม่สำเร็จ',
          });
        }
      });
  }

  return (
    <Page theme={pageTheme.home}>
      <Header style={HeaderCustom} title={`Borrow Book`}>
        <Avatar alt="Remy Sharp" src="../../image/account.jpg" />
        <div style={{ marginLeft: 10 }}>Manuschanok Srikhrueadong</div>
      </Header>
      <Content>
        <Container maxWidth="sm">
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>
            <Grid item xs={3}>
              <div className="{classes.paper}">หนังสือ</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกหนังสือ</InputLabel>
                <Select
                  name="book"
                  value={bookborrow.book || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {books.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.bookName}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>ผู้ยืม</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกผู้ยืม</InputLabel>
                <Select
                  name="user"
                  value={bookborrow.user || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {users.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.uSERNAME}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>จุดบริการ</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกจุดบริการ</InputLabel>
                <Select
                  name="servicepoint"
                  value={bookborrow.servicepoint || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {servicepoints.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.cOUNTERNUMBER}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>


            <Grid item xs={3}>
              <div className={classes.paper}>เวลา</div>
            </Grid>
            <Grid item xs={9}>
              <form className={classes.container} noValidate>
                <TextField
                  label="เลือกเวลา"
                  name="added"
                  type="date"
                  value={bookborrow.added || ''} // (undefined || '') = ''
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  onChange={handleChange}
                />
              </form>
            </Grid>

            <Grid item xs={3}></Grid>
            <Grid item xs={9}>
              <Button
                variant="contained"
                color="primary"
                size="large"
                startIcon={<SaveIcon />}
                onClick={save}
              >
                บันทึกการดู
              </Button>
            </Grid>
          </Grid>
        </Container>
      </Content>
    </Page>
  );
};

export default Bookborrow;
