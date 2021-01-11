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
  EntBook,
  EntServicePoint,
  EntUser,
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
const username = { givenuser: 'Manuschanok Srikhrueadong' };
export default function Create() {
  const profile = { givenName: 'ยินดีต้อนรับสู่ ระบบห้องสมุด' };
  const classes = useStyles();
  const api = new DefaultApi();
  const [loading, setLoading] = useState(true);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [books, setBooks] = useState<EntBook[]>(Array);
  const [users, setUsers] = useState<EntUser[]>(Array);
  const [servicepoints, setServicepoints] = useState<EntServicePoint[]>(Array);
  useEffect(() => {
    const getUser = async () => {
      const res = await api.listUser({ limit: 10, offset: 0 });
      setLoading(false);
      setUsers(res);
    };
    getUser();

    const getBooks = async () => {
      const res = await api.listBook({ limit: 10, offset: 0 });
      setLoading(false);
      setBooks(res);
    };
    getBooks();

    const getServicepoints = async () => {
      const res = await api.listServicepoint({ limit: 10, offset: 0 });
      setLoading(false);
      setServicepoints(res);
    };
    getServicepoints();
  }, [loading]);


  const ServicepointIDhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setServicepointID(event.target.value as number);
  };
  const bookIDhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setBookID(event.target.value as number);
  };
  const UserIDhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setUserID(event.target.value as number);
  };

 
  const [bookID, setBookID] = useState(Number);
  const [servicePointID, setServicepointID] = useState(Number);
  const [userID, setUserID] = useState(Number);
  const createBookBorrow = async () => {
    const bookBorrow = {
      bookID : bookID,
      servicePointID: servicePointID,
      userID : userID,
    };
    console.log(bookBorrow);
    const res: any = await api.createBookborrow({ bookborrow: bookBorrow });
    setStatus(true);
    if (res.id != '') {
      setAlert(true);
    } else {
      setAlert(false);
    }
  };
  return (
    <Page theme={pageTheme.home}>
      <Header
        title={"ระบบยืมหนังสือ"}
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
        <ContentHeader title="เพิ่มข้อมูลการยืมหนังสือ">
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
              <tr><td>หนังสือ</td><td>
                <FormControl
                  className={classes.margin}
                  variant="outlined"
                >
                  <InputLabel id="book-label"><font size='5'>หนังสือ</font></InputLabel>
                  <Select
                    labelId="book-label"
                    id="book"
                    value={bookID}
                    onChange={bookIDhandleChange}
                    style={{ width: 400, height: '7vh' }}
                  >
                    {books.map((item: EntBook) => (
                      <MenuItem value={item.id}>{item.bookName}</MenuItem>
                    ))}
                  </Select>
                </FormControl>

              </td>
              </tr>
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
              <tr><td>
                จุดบริการ
            </td><td>
                  <FormControl
                    className={classes.margin}
                    variant="outlined"
                  >
                    <InputLabel id="servicepoint"><font size='5'>จุดบริการ</font></InputLabel>
                    <Select
                      labelId="servicepoint"
                      id="servicepoint"
                      value={servicePointID}
                      onChange={ServicepointIDhandleChange}
                      style={{ width: 400, height: '7vh' }}
                    >
                      {servicepoints.map((item: EntServicePoint) => (
                        <MenuItem value={item.id}>{item.cOUNTERNUMBER}</MenuItem>
                      ))}
                    </Select>
                  </FormControl>
                </td>
              </tr>

              <td></td><td>
                <div className={classes.margin}>
                  <Button
                    onClick={() => {
                      createBookBorrow();
                    }}
                    variant="contained"
                    color="primary"
                  >
                    [บันทึกการยืม]
           </Button>
                  <Button
                    style={{ marginLeft: 20 }}
                    component={RouterLink}
                    to="/welcome"
                    variant="contained"
                  >
                    กลับ
           </Button>
                </div>
          </form></td></table>
        </div>

      </Content>
    </Page>
  );
}
