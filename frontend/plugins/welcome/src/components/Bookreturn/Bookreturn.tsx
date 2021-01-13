import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
  Content, Header, Page, pageTheme, ContentHeader,
} from '@backstage/core';
import {
  Typography, TextField, Button, makeStyles,
  Theme, FormControl, InputLabel, MenuItem,
  Select, createStyles,
} from '@material-ui/core';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';
import {
  EntUser,
  EntBookborrow,
  EntLocation,
} from '../../api/models/';
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
      width: '50ch',
    },
  }),
);
const username = { givenuser: 'Ton' };
export default function Create() {
  const profile = { givenName: 'ยินดีต้อนรับสู่ ระบบห้องสมุด' };
  const classes = useStyles();
  const api = new DefaultApi();
  const [loading, setLoading] = useState(true);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
<<<<<<< HEAD
  const [iduser,setIduser] = React.useState<number>(0)
=======
  
>>>>>>> 7a01a0a75785e4aa3a74f4bf2b3c889adcb492c7
  const [users, setUsers] = React.useState<EntUser[]>(Array);
  const [bookborrows, setBookborrows] = React.useState<EntBookborrow[]>([]);
  const [locations, setLocations] = React.useState<EntLocation[]>(Array);
  useEffect(() => {
    const getUser = async () => {
<<<<<<< HEAD
      const res = await api.listUser({ limit: 10, offset: 0 });
      setLoading(false);
      setUsers(res);
    };
    getUser();

    const getBookborrows = async () => {
      const res = await api.getBookborrowuser({id:1});
      setLoading(false);
      setBookborrows(res);
      console.log("bookborrows => "+bookborrows);
    };
    getBookborrows();

=======
      const res = await api.listUser();
      setLoading(false);
      setUsers(res);
      console.log("users => "+users);
    };
    getUser();

>>>>>>> 7a01a0a75785e4aa3a74f4bf2b3c889adcb492c7
    const getLocations = async () => {
      const res = await api.listLocation({ limit: 10, offset: 0 });
      setLoading(false);
      setLocations(res);
    };
    getLocations();
<<<<<<< HEAD
=======
    
>>>>>>> 7a01a0a75785e4aa3a74f4bf2b3c889adcb492c7
  }, [loading]);

  const UserIDhandleChange = (event: React.ChangeEvent<{ value: any }>) => {
    setUserID(event.target.value as number);
<<<<<<< HEAD
    setIduser(event.target.value as number)
  };
  const BookborrowIDhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setBookborrowID(event.target.value as number);
=======
    const getBookborrows = async () => {
      const res = await api.getBookborrowuser({id:event.target.value as number});
      setBookborrows(res);
      console.log("bookborrows => "+bookborrows);
    };
    getBookborrows();
  };
  const BookborrowIDhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setBookborrowID(event.target.value as number);
  };

  const LocationIDhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setLocationID(event.target.value as number);
>>>>>>> 7a01a0a75785e4aa3a74f4bf2b3c889adcb492c7
  };
  

<<<<<<< HEAD
  const LocationIDhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setLocationID(event.target.value as number);
  };


  const [userID, setUserID] = useState(Number);
  const [bookborrowID, setBookborrowID] = useState(Number);
  const [locationID, setLocationID] = useState(Number);
  
  const createBookreturn = async () => {
    const bookreturn = {
      userID : userID,
      bookborrowID : bookborrowID,
      locationID: locationID,
    };
    console.log(bookreturn);
    const res: any = await api.createBookreturn({ bookreturn: bookreturn });
    setStatus(true);
    if (res.id != '') {
      setAlert(true);
    } else {
      setAlert(false);
    }
  };
=======
  const [userID, setUserID] = useState(Number);
  const [bookborrowID, setBookborrowID] = useState(Number);
  const [locationID, setLocationID] = useState(Number);
  
  const createBookreturn = async () => {
    if ( (userID != "") && (bookborrowID != "") && (locationID != "") ){
    const bookreturn = {
      userID : userID,
      bookborrowID : bookborrowID,
      locationID: locationID,
    };
    console.log(bookreturn);
    const res: any = await api.createBookreturn({ bookreturn: bookreturn });
    setStatus(true);
    if (res.id != '') {
      setAlert(true);
      window.location.reload(false);
    }
    } else {
      setAlert(false);
      setStatus(true);
    }
    const timer = setTimeout(() => {
      setStatus(false);
  }, 3000);
  };
>>>>>>> 7a01a0a75785e4aa3a74f4bf2b3c889adcb492c7
  return (
    <Page theme={pageTheme.home}>
      <Header
        title={"ระบบคืนหนังสือ"}
      >
        <Button
          component={RouterLink}
          to="/login"
          variant="contained"
          color="primary">
          ออกจากระบบ
        </Button>
      </Header>
      <Content>
        <ContentHeader title="เพิ่มข้อมูลการคืนหนังสือ">
          <div>

          </div>
          {status ? (
            <div>
              {alert ? (
                <Alert severity="success">
                  This is a success alert — check it out!
                </Alert>
              ) : (
                  <Alert severity="warning" style={{ marginTop: 20 }}>
                    This is a warning alert — check it out!
                  </Alert>
                )}
            </div>
          ) : null}
        </ContentHeader>


        <div className={classes.root}>
          <form noValidate autoComplete="off">
            <table align="center">

              <tr><td>
                สมาชิกห้องสมุด
            </td><td>
                  <div>
                    <FormControl
                      className={classes.margin}
                      variant="outlined"
                    >
                      <InputLabel id="user-label"><font size='5'>สมาชิกห้องสมุด</font></InputLabel>
                      <Select
                        labelId="user-label"
                        id="user"
                        value={userID}
                        onChange={UserIDhandleChange}
                        style={{ width: 400, height: '7vh' }}
                      >
                        {users.map((item: EntUser) => (
                          <MenuItem value={item.id}>{item.uSERNAME}</MenuItem>
                        ))}
                      </Select>
                    </FormControl>
                  </div>
                </td>
              </tr>

              <tr><td>รายการยืมหนังสือ</td><td>
                <FormControl
                  className={classes.margin}
                  variant="outlined"
                >
                  <InputLabel id="bookborrow-label"><font size='5'>หนังสือ</font></InputLabel>
                  <Select
                    labelId="bookborrow-label"
                    id="bookborrow"
                    value={bookborrowID}
                    onChange={BookborrowIDhandleChange}
                    style={{ width: 400, height: '7vh' }}
                  >
                    {bookborrows.map(item => {
                    return (
                      <MenuItem value={item.id}>{item.id}</MenuItem>
                    );})}
                  </Select>
                </FormControl>

              </td>
              </tr>

              <tr><td>
                สถานที่คืนหนังสือ
            </td><td>
                  <FormControl
                    className={classes.margin}
                    variant="outlined"
                  >
                    <InputLabel id="location-label"><font size='5'>สถานที่คืนหนังสือ</font></InputLabel>
                    <Select
                      labelId="location-label"
                      id="location"
                      value={locationID}
                      onChange={LocationIDhandleChange}
                      style={{ width: 400, height: '7vh' }}
                    >
                      {locations.map((item: EntLocation) => (
                        <MenuItem value={item.id}>{item.lOCATIONNAME}</MenuItem>
                      ))}
                    </Select>
                  </FormControl>
                </td>
              </tr>

              <td></td><td>
                <div className={classes.margin}>
                  <Button
                    onClick={() => {
                      createBookreturn();
                    }}
                    variant="contained"
                    color="primary"
                  >
                    บันทึกการยืม
           </Button>
                  <Button
                    style={{ marginLeft: 20 }}
                    component={RouterLink}
                    to="/welcome"
                    variant="contained"
                  >
                    กลับ
           </Button>
           </div></td>
           </table></form>
        </div>

      </Content>
    </Page>
  );
}