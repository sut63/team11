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
  Avatar,
  Button,
} from '@material-ui/core';
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command
import {
    EntUser,
    EntBookborrow,
    EntLocation,
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

interface Bookreturn {
  user: number;
  bookborrow: number;
  location: number;
  deadline: Date;
  // create_by: number;
}

const Bookreturn: FC<{}> = () => {
  const classes = useStyles();
  const http = new DefaultApi();

  const [bookreturn, setBookreturn] = React.useState<
    Partial<Bookreturn>
  >({});

  const [users, setUsers] = React.useState<EntUser[]>([]);
  const [bookborrows, setBookborrows] = React.useState<EntBookborrow[]>([]);
  const [locations, setLocations] = React.useState<EntLocation[]>([]);
  

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

  const getBookborrows = async () => {
    const res = await http.listBookborrow({ limit: 10, offset: 0 });
    setBookborrows(res);
  };

  const getLocations = async () => {
    const res = await http.listLocation({ limit: 10, offset: 0 });
    setLocations(res);
  };



  // Lifecycle Hooks
  useEffect(() => {
    getUsers();
    getBookborrows();
    getLocations();
  }, []);

  // set data to object playlist_video
  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>,
  ) => {
    const name = event.target.name as keyof typeof Bookreturn;
    const { value } = event.target;
    setBookreturn({ ...bookreturn, [name]: value });
    console.log(bookreturn);
  };

  // clear input form
  function clear() {
    setBookreturn({});
  }

  // function save data
  function save() {
    const apiUrl = 'http://localhost:8080/api/v1/bookreturn';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(bookreturn),
    };

    console.log(bookreturn); // log ดูข้อมูล สามารถ Inspect ดูข้อมูลได้ F12 เลือก Tab Console

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
      <Header style={HeaderCustom} title={`Return Book`}>
        <Avatar alt="Remy Sharp" src="../../image/account.jpg" />
        <div style={{ marginLeft: 10 }}>Ton</div>
      </Header>
      <Content>
        <Container maxWidth="sm">
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>

            <Grid item xs={3}>
              <div className={classes.paper}>ผู้คืน</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกผู้คืน</InputLabel>
                <Select
                  name="user"
                  value={bookreturn.user || ''} // (undefined || '') = ''
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
              <div className="{classes.paper}">รายการยืมหนังสือ</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกรายการ</InputLabel>
                <Select
                  name="bookborrow"
                  value={bookreturn.bookborrow || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {bookborrows.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.bookID}
                      </MenuItem>// I think it incorrect**
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            

            <Grid item xs={3}>
              <div className={classes.paper}> สถานที่คืน</div>
            </Grid>
            <Grid item xs={9}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกสถานที่คืน</InputLabel>
                <Select
                  name="location"
                  value={bookreturn.location || ''} // (undefined || '') = ''
                  onChange={handleChange}
                >
                  {locations.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.lOCATIONNAME}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
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
                บันทึกการคืน
              </Button>
            </Grid>
          </Grid>
        </Container>
      </Content>
    </Page>
  );
};

export default Bookreturn;
