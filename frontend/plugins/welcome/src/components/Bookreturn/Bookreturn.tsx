import React, { useState, useEffect } from 'react';
//import SaveAltIcon from '@material-ui/icons/SaveAlt';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
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
export default function Create() {
  const classes = useStyles();
  const api = new DefaultApi();
  const [loading, setLoading] = useState(true);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  
  const [users, setUsers] = React.useState<EntUser[]>(Array);
  const [bookborrows, setBookborrows] = React.useState<EntBookborrow[]>([]);
  const [locations, setLocations] = React.useState<EntLocation[]>(Array);

  const idString = JSON.parse(String(localStorage.getItem("userID")));
  const idInt = parseInt(idString);
  useEffect(() => {
    const getUser = async () => {
      const res = await api.listUser();
      setLoading(false);
      setUsers(res);
      console.log("users => "+users);
    };
    getUser();

    setUserID(idInt);

    const getBookborrows = async () => {
      const res = await api.getBookborrowuser({id:idInt});
      setBookborrows(res);
      console.log("bookborrows => "+bookborrows);
    };
    getBookborrows();

    const getLocations = async () => {
      const res = await api.listLocation({ limit: 10, offset: 0 });
      setLoading(false);
      setLocations(res);
    };
    getLocations();
    
  }, [loading]);

  /*const UserIDhandleChange = (event: React.ChangeEvent<{ value: any }>) => {
    setUserID(event.target.value as number);
  };*/
  const BookborrowIDhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setBookborrowID(event.target.value as number);
  };

  const LocationIDhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setLocationID(event.target.value as number);
  };
  

  const [userID, setUserID] = useState(Number);
  const [bookborrowID, setBookborrowID] = useState(Number);
  const [locationID, setLocationID] = useState(Number);
  
  const createBookreturn = async () => {
    if ( (userID != 0) && (bookborrowID != 0) && (locationID != 0) ){
    const bookreturn = {
      userID : userID,
      bookborrowID : bookborrowID,
      locationID: locationID,
    };
    console.log(bookreturn);
    const res: any = await api.createBookreturn({ bookreturn: bookreturn });
    setStatus(true);
    if (res.id != '') {
      //await api.deleteBookborrow({id:bookborrowID})
      setAlert(true);
    }
    } else {
      setAlert(false);
      setStatus(true);
    }
    const timer = setTimeout(() => {
      setStatus(false);
      window.location.reload(false);
  }, 5000);
  };

  const resetLocalStorage = async () => {
    localStorage.setItem("userID", JSON.stringify(null))
    localStorage.setItem("role", JSON.stringify(null))
    localStorage.setItem("valid", JSON.stringify(null))
    localStorage.setItem("userName", JSON.stringify(null))
    window.location.href = "/"
  }

  return (
    <Page theme={pageTheme.home}>
      <Header
        title={"ระบบคืนหนังสือ"}
      >

      </Header>
      <Content>
        <ContentHeader title="เพิ่มข้อมูลการคืนหนังสือ">
        <Button
            // disabled={LogoutBtn}
            variant="contained"
            color="primary"
            className={classes.button}
            startIcon={<LockOutlinedIcon />}
            onClick={() => {
              resetLocalStorage();
            }}>
            ล็อกเอ้าท์
          </Button>
          
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
                        //onChange={UserIDhandleChange}
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
                  <InputLabel id="bookborrow-label"><font size='5'>รายการยืมหนังสือ</font></InputLabel>
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
                    บันทึกการคืน
           </Button>
           
                  <Button
                    
                    style={{ marginLeft: 20 }}
                    component={RouterLink}
                    to="/"
                    variant="contained"
                  >
                  
                    กลับ
           </Button>

           {status ? (
            <div>
              {alert ? (
                <Alert severity="success">
                   บันทึกการคืนสำเร็จ
                </Alert>
              ) : (
                  <Alert severity="warning" style={{ marginTop: 20 }}>
                    บันทึกการคืนไม่สำเร็จ
                  </Alert>
                )}
            </div>
          ) : null}
           </div></td>
           </table></form>
        </div>

      </Content>
    </Page>
  );
}