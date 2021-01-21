import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import SaveAltIcon from '@material-ui/icons/SaveAlt';


import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
} from '@backstage/core';
import AccountCircle from '@material-ui/icons/AccountCircle';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert, AlertTitle } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';
//เรียกentมาใช้
import {
  EntAuthor,
  EntResearchtype,
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
    button: {
      margin: theme.spacing(1),
    },
    selectEmpty: {
      marginTop: theme.spacing(2),
    },
  }),
);

interface researchField {
  /**
     * 
     * @type {string}
     * @memberof researchField
     */
    docname?: string;
  /**
     * 
     * @type {number}
     * @memberof researchField
     */
    pagenumber?: number;
  /**
     * 
     * @type {number}
     * @memberof researchField
     */
    yearnumber?: number;
}

export default function Create() {
  const classes = useStyles();
  const name = JSON.parse(String(localStorage.getItem("userName")));
  const userName = name
  const profile = { thisName: 'ระบบงานวิจัยและฐานข้อมูลออนไลน์' };
  const api = new DefaultApi();
  //-------
  const [users, setUsers] = useState<EntUser[]>([]);
  const [authors, setAuthors] = useState<EntAuthor[]>([]);
  const [researchtypes, setResearchtypes] = useState<EntResearchtype[]>([]);

  const [userid, setUser] = useState(Number);
  const [authorid, setAuthor] = useState(Number);
  const [researchtypeid, setResearchtype] = useState(Number);

  const [datetime, setdatetime] = useState(String);
  const [researchField, setResearch] = useState<Partial<researchField>>({});
  const [loading, setLoading] = useState(true);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);

  const idString = JSON.parse(String(localStorage.getItem("userID")));
  const idInt = parseInt(idString);
  //function
  useEffect(() => {
    const getUser = async () => {
      const res = await api.listUser();
      setLoading(false);
      setUsers(res);
    };
    getUser();

    setUser(idInt);

    const getAuthor = async () => {
      const res = await api.listAuthor({ offset: 0 });
      setLoading(false);
      setAuthors(res);
      console.log(res);
    };
    getAuthor();

    const getResearchtype = async () => {
      const res = await api.listResearchtype({ offset: 0 });
      setLoading(false);
      setResearchtypes(res);
      console.log(res);
    };
    getResearchtype();

  }, [loading]);

  const [docNameError, setDOCNAMEError] = useState('');
  const [pageNumberError, setPAGENUMBERError] = useState('');
  const [yearNumberError, setYEARNUMBERError] = useState('');
  const [alertMessage, setAlertMessage] = useState('');

  const handleChange = (event: React.ChangeEvent<{ id?: string; value: any }>) => {
    const id = event.target.id as keyof typeof Create;
    const { value } = event.target;
    const validateValue = value.toString()
    checkPattern(id, validateValue)
    setResearch({ ...researchField, [id]: value });
  };

  const validateDOCNAME = (val: string) => {
    return val.match("^[0-9a-zA-Zก-๙]+$");
  }
  const validatePAGENUMBER = (val: number) => {
    return val < 1 ? false : true
  }
  const validateYEARNUMBERE = (val: number) => {
    return val > 2999  ? false : true && val < 1000  ? false : true
  }

  const checkPattern = (id: string, value: any) => {
    console.log(value);
    switch (id) {
      case 'docname':
        validateDOCNAME(value) ? setDOCNAMEError('') : setDOCNAMEError('กรุณากรอกชื่อเรื่อง (ชื่อเรื่องห้ามมีตัวอักษรพิเศษ)');
        return;
      case 'pagenumber':
        validatePAGENUMBER(value) ? setPAGENUMBERError('') : setPAGENUMBERError('กรุณากรอกจำนวนหน้า (ห้ามเป็น 0)');
        return;
      case 'yearnumber':
        validateYEARNUMBERE(value) ? setYEARNUMBERError('') : setYEARNUMBERError('กรุณากรอกปีที่พิมพ์ (ให้มากกว่าหรือเท่ากับ 1000 เเต่ไม่เกิน 3000 ปีพ.ศ.)');
        return;
      default:
        return;
    }
  }

  const checkCaseSaveError = (field: string) => {
    switch(field) {
      case 'DOC_NAME':
        setAlertMessage("error ข้อมูล field docname ผิด");
        return;
      case 'PAGE_NUMBER':
        setAlertMessage("error ข้อมูล field pagenumber ผิด");
        return;
      case 'YEAR_NUMBER':
        setAlertMessage("error ข้อมูล field yearnumber ผิด");
        return;
      default:
        setAlertMessage("บันทึกข้อมูลไม่สำเร็จ");
        return;
    }
  }




  const handledatetimeChange = (event: any) => {
    setdatetime(event.target.value as string);
  };

  const handleAuthorchange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setAuthor(event.target.value as number);
  };

  const handleResearchtypechange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setResearchtype(event.target.value as number);
  };


  const createResearch = async () => {
    
          const research = {
        register: userid,
        myDoc: authorid,
        docType: researchtypeid,
        docname: String(researchField.docname),
        pagenumber:Number(researchField.pagenumber),
        yearnumber:Number(researchField.yearnumber),
        date: datetime + ":00+07:00",
      };

      const apiUrl = 'http://localhost:8080/api/v1/researches';
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(research),
    };

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        if (data.status === true) {
          setStatus(true);
          setAlert(true);
        } else {
          setStatus(true);
          setAlert(false);
          checkCaseSaveError(data.error.Name)
        }
      });
      const timer = setTimeout(() => {
        setStatus(false);
        //window.location.reload(false);
      }, 1000000000);
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
        title={`${profile.thisName}`}
        subtitle="Online research system and database"
      >
        <Button
          variant="contained"
          color="secondary"
          className={classes.button}
          startIcon={<LockOutlinedIcon />}
          onClick={() => {
            resetLocalStorage();
          }}>
          ล็อกเอ้าท์
          </Button>

      </Header>

      <Content>
        <ContentHeader title="เพิ่มข้อมูลงานวิจัยและฐานข้อมูลออนไลน์">        </ContentHeader>

        <div><center>
          <AccountCircle fontSize="large" className={classes.large} /><br></br>
          <Button variant="contained" color='primary' size='large' >
            <font size='3'>helo are you 🌵 {userName}🌵</font>
          </Button>
        </center>
        </div>

        <div className={classes.root}>
          <form noValidate autoComplete="off">

          <div>
            <FormControl required className={classes.formControl}>
                <InputLabel id="demo-simple-select-required-label">Library</InputLabel>
                <Select
                  disabled={true}
                  labelId="demo-simple-select-required-label"
                  id="demo-simple-select-required"
                  value={userid || ''}
                  className={classes.selectEmpty}
                  style={{ width: 200, marginLeft: 30, marginRight: -10 }}
                >
                  {users.map((item: EntUser) => (
                    <MenuItem value={item.id}>{item.uSEREMAIL}</MenuItem>
                  ))}
                </Select>
              </FormControl>
            </div>

            <div className={classes.paper}><strong>ชื่องานวิจัย</strong></div>
            <FormControl required className={classes.formControl}>
                  <TextField style={{ width: 300, marginLeft: 30, marginRight: -10 }} error={docNameError ? true : false} id="docname"
                    helperText={docNameError} onChange={handleChange} label="ชื่องานวิจัย*"
                    value={researchField.docname || ''} />
                </FormControl>
            

            <div className={classes.paper}><strong>จำนวนหน้า</strong></div>
            <form className={classes.root} noValidate autoComplete="off">
            <FormControl required className={classes.formControl}>
                  <TextField style={{ width: 150 }} error={pageNumberError ? true : false} id="pagenumber"
                    helperText={pageNumberError} type="number" InputProps={{ inputProps: { min: 1, max: 6 } }}
                    onChange={handleChange} label="จำนวนหน้า"
                    value={researchField.pagenumber} />
                </FormControl>
            </form>

            <div className={classes.paper}><strong>ปีที่พิมพ์</strong></div>
            <form className={classes.root} noValidate autoComplete="off">
            <FormControl required className={classes.formControl}>
                  <TextField style={{ width: 150 }} error={yearNumberError ? true : false} id="yearnumber"
                    helperText={yearNumberError} type="number" InputProps={{ inputProps: { min: 1, max: 6 } }}
                    onChange={handleChange} label="ปีที่พิมพ์ (พ.ศ.)"
                    value={researchField.yearnumber} />
                </FormControl>
            </form>
            
            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <div className={classes.paper}><strong>ชื่อผู้แต่ง</strong></div>
                <InputLabel id="author-label"></InputLabel>
                <Select
                  labelId="author-label"
                  id="ผู้เเต่ง"
                  value={authorid}
                  onChange={handleAuthorchange}
                  style={{ width: 400 }}
                >
                  {authors.map((item: EntAuthor) => (
                    <MenuItem value={item.id}>{item.name}</MenuItem>
                  ))}
                </Select>
              </FormControl>
            </div>

            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <div className={classes.paper}><strong>ชื่อประเภทงานวิจัย</strong></div>
                <InputLabel id="researchtype-label"></InputLabel>
                <Select
                  labelId="researchtype-label"
                  id="ประเภทงานวิจัย"
                  value={researchtypeid}
                  onChange={handleResearchtypechange}
                  style={{ width: 400 }}
                >
                  {researchtypes.map((item: EntResearchtype) => (
                    <MenuItem value={item.id}>{item.tYPENAME}</MenuItem>
                  ))}
                </Select>
              </FormControl>
            </div>

            <FormControl className={classes.margin} >
              <TextField
                id="Datetime"
                label="วัน-เวลา"
                type="datetime-local"
                value={datetime}
                onChange={handledatetimeChange}
                className={classes.textField}
                InputLabelProps={{
                  shrink: true,
                }}
                style={{ width: 250 }}

              />
            </FormControl>

            <div className={classes.margin}>
              <Button
                onClick={() => {
                  createResearch();
                }}
                variant="contained"
                color="primary"
              >
                ยืนยันการบันทึก
             </Button>
             <div>
             {status ? (
            <div>
              {alert ? (
                <Alert severity="success">
                  เพิ่มบันทึกเรียบร้อย!
                </Alert>
              ) : (
                  <Alert severity="warning" style={{ marginTop: 20 }}>
                    {alertMessage}
                  </Alert>
                )}
            </div>
          ) : null}
                    </div>
            </div>
          </form>
        </div>
      </Content>
    </Page>
  );
}