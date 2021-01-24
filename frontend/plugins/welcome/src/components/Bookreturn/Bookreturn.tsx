import React, { useState, useEffect } from 'react';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
import { Link as RouterLink } from 'react-router-dom';
import Swal from 'sweetalert2';
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

export default function Create() {
  const classes = useStyles();
  const api = new DefaultApi();
  const [loading, setLoading] = useState(true);
  // const [status, setStatus] = useState(false);
  // const [alert, setAlert] = useState(true);

  const [users, setUsers] = React.useState<EntUser[]>(Array);
  const [bookborrows, setBookborrows] = React.useState<EntBookborrow[]>([]);
  const [locations, setLocations] = React.useState<EntLocation[]>(Array);

  const idString = JSON.parse(String(localStorage.getItem("userID")));
  const idInt = parseInt(idString);

  const [damagedpointError, setdamagedpointError] = React.useState('');
  const [damagedpointnameError, setDamagedpointnameError] = React.useState('');
  const [lostError, setLostError] = React.useState('');

  const [damagedpoint, setDamagedpoint] = useState(Number);
  const [damagedpointname, setDamagedpointname] = useState(String);
  const [lost, setLost] = useState(String);

  const [errordamagedpoint, setErrorDamagedpoint] = useState(true);
  const [errordamagedpointname, setErrorDamagedpointName] = useState(true);
  const [errorlost, setErrorLost] = useState(true);



  // function validate DamagedPoint
  const ValidateDamagedPoint = (val: number) => {
    val < 1 || val > 10 ? setErrorDamagedpoint(false) : setErrorDamagedpoint(true);
  }

  // function validate DamagedPointName
  const ValidateDamagedPointName = (val: string) => {
    val.match("^[a-zA-Z]+$") ? setErrorDamagedpointName(true) : setErrorDamagedpointName(false);
  }

  // function validate Lost
  const ValidateLost = (lost: string) => {
    lost.match("^[a-zA-Z]+$") ? setErrorLost(true) : setErrorLost(false);
  }

  // สำหรับตรวจสอบรูปแบบข้อมูลที่กรอก ว่าเป็นไปตามที่กำหนดหรือไม่
  /*
  const checkPattern = (id: string, value: string) => {
    switch (id) {
      case 'damaged_point':
        DamagedPoint(Number(value)) ? setdamagedpointError('') : setdamagedpointError('ต้องเป็นตัวเลข 1-10 ');
        return;
      case 'damaged_point_name':
        DamagedPointName(value) ? setDamagedpointnameError('') : setDamagedpointnameError('จุดที่เสียหายเป็นภาษาอังกฤษเท่านั้น เช่น TopFront,BottomBack');
        return;
      case 'lost':
        Lost(value) ? setLostError('') : setLostError('ถ้าหายให้พิมพ์ lost ถ้าไม่พิมพ์ no')
        return;
      default:
        return;
    }
  }
  */

  const alertMessage = (icon: any, title: any) => {
    Toast.fire({
      icon: icon,
      title: title,
    });
  }

  const checkCaseSaveError = (field: string) => {
    switch (field) {
      case 'DAMAGED_POINT':
        alertMessage("error", "ต้องเป็นตัวเลข 1-10 เท่านั้น");
        return;
      case 'DAMAGED_POINTNAME':
        alertMessage("error", "จุดที่เสียหายเป็นภาษาอังกฤษเท่านั้น เช่น TopFront,BottomBack");
        return;
      case 'LOST':
        alertMessage("error", "ถ้าหายให้พิมพ์ lost ถ้าไม่พิมพ์ no");
        return;
      default:
        alertMessage("error", "บันทึกข้อมูลไม่สำเร็จ");
        return;
    }
  }


  useEffect(() => {
    const getUser = async () => {
      const res = await api.listUser();
      setLoading(false);
      setUsers(res);
      console.log("users => " + users);
    };
    getUser();

    setUserID(idInt);

    const getBookborrows = async () => {
      const res = await api.getBookborrowuser({ id: idInt });
      setBookborrows(res);
      console.log("bookborrows => " + bookborrows);
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

  const DamagedPointhandleChange = (event: React.ChangeEvent<{ value: any }>) => {
    setDamagedpoint(event.target.value as number);
    ValidateDamagedPoint(event.target.value as number);
  };

  const DamagedPointNamehandleChange = (event: React.ChangeEvent<{ value: any }>) => {
    setDamagedpointname(event.target.value as string);
    ValidateDamagedPointName(event.target.value as string);
  };

  const LosthandleChange = (event: React.ChangeEvent<{ value: any }>) => {
    setLost(event.target.value as string);
    ValidateLost(event.target.value as string);
  };

  /*const handleChange = (event: React.ChangeEvent<{ id?: string; value: any }>) => {
    const id = event.target.id as keyof typeof Create;
    const { value } = event.target;
    const validateValue = value.toString()
    checkPattern(id, validateValue)
    setBookreturn({ ...returnbook, [id]: value });
  };*/

  const [userID, setUserID] = useState(Number);
  const [bookborrowID, setBookborrowID] = useState(Number);
  const [locationID, setLocationID] = useState(Number);


  const createBookreturn = async () => {
    const apiUrl = 'http://localhost:8080/api/v1/bookreturns';

    const bookreturn = {
      userID: userID,
      bookborrowID: bookborrowID,
      locationID: locationID,
      damagedPoint: Number(damagedpoint),
      damagedPointName: String(damagedpointname),
      lost: String(lost),

    };
    console.log(bookreturn);

    /*const ErrorCaseCheck = (casename: string) => {
      if (casename == "wheel_center") { setErrorMessege("ค่า ศูนย์ล้อ ต้องมากกว่า 0"); }
      else if (casename == "sound_level") { setErrorMessege("ค่า ระดับเสียง ต้องมากกว่า 0"); }
      else if (casename == "blacksmoke") { setErrorMessege("ค่า ควันดำ ต้องอยู่ในช่วง 1-100"); }
      else { setErrorMessege("บันทึกไม่สำเร็จ"); }
    }*/

    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(bookreturn),
    };

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.status === true) {
          Toast.fire({
            icon: 'success',
            title: 'บันทึกข้อมูลสำเร็จ',
          });
        } else {
          //ErrorCaseCheck(data.error.Name)
          checkCaseSaveError(data.error.Name)
        }
      });
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
                        disabled={true}
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
                      );
                    })}
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

              <div>
                <FormControl
                  //fullWidth
                  className={classes.margin}
                  variant="outlined"
                >
                  <TextField
                    error={errordamagedpoint ? false : true}
                    id="DAMAGED_POINT"
                    label="ตำแหน่งที่ชำรุด"
                    variant="outlined"
                    type="number"
                    size="medium"
                    helperText={errordamagedpoint ? "" : "ต้องเป็นตัวเลข 0-10 เท่านั้น"}
                    InputProps={{ inputProps: { min: 1, max: 10 } }}
                    value={damagedpoint}
                    onChange={DamagedPointhandleChange}
                  />
                </FormControl>
              </div>

              <div>
                <FormControl
                  //fullWidth
                  className={classes.margin}
                  variant="outlined"
                >
                  <TextField
                    error={errordamagedpointname ? false : true}
                    id="DAMAGED_POINTNAME"
                    label="ชื่อตำแหน่งที่ชำรุด"
                    variant="outlined"
                    type="string"
                    size="medium"
                    helperText={errordamagedpointname ? "" : "จุดที่เสียหายเป็นภาษาอังกฤษเท่านั้น เช่น TopFront,BottomBack"}
                    value={damagedpointname}
                    onChange={DamagedPointNamehandleChange}
                  />
                </FormControl>
              </div>

              <div>
                <FormControl
                  //fullWidth
                  className={classes.margin}
                  variant="outlined"
                >
                  <TextField
                    error={errorlost ? false : true}
                    id="LOST"
                    label="กรณีหนังสือหาย"
                    variant="outlined"
                    type="string"
                    size="medium"
                    helperText={errorlost ? "" : "ถ้าหายให้พิมพ์ lost ถ้าไม่ พิมพ์ no"}
                    value={lost}
                    onChange={LosthandleChange}
                  />
                </FormControl>
              </div>

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


                </div></td>
            </table>
          </form>
        </div>

      </Content>
    </Page>
  );
}