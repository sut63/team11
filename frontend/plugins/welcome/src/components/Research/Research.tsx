import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
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
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';
import LockOutlinedIcon from '@material-ui/icons/LockOutlined';

//เรียกentมาใช้
import { EntUser } from '../../api/models/EntUser';
import { EntAuthor } from '../../api/models/EntAuthor'; 
import { EntResearchtype } from '../../api/models/EntResearchtype'; 

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
  }),
);

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
  const [loading, setLoading] = useState(true);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);

  const [title, setTitle] = useState(String);
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

  const handleTitleChange = (event: any) => {
    setTitle(event.target.value as string);
  };

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
    if ((title != null) && (title != "") &&(datetime != null) && (datetime != "") && (authorid != null) && (researchtypeid != null)) {
      const research = {
        register: userid,
        myDoc: authorid,
        docType: researchtypeid,
        docname: title,
        date: datetime + ":00+07:00",
      };
      console.log(research);
      const res: any = await api.createResearch({ research: research });
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
        <ContentHeader title="เพิ่มข้อมูลงานวิจัยและฐานข้อมูลออนไลน์">

          {status ? (
            <div>
              {alert ? (
                <Alert severity="success">
                  เพิ่มบันทึกเรียบร้อย!
                </Alert>
              ) : (
                  <Alert severity="warning" style={{ marginTop: 20 }}>
                    บันทึกไม่สำเร็จกรอกข้อมูลให้ครบ
                  </Alert>
                )}
            </div>
          ) : null}
        </ContentHeader>

        <div><center>
          <AccountCircle fontSize="large" className={classes.large} /><br></br>
          <Button variant="contained" color='primary' size='large' >
            <font size='3'>helo are you 🌵 {userName}🌵</font>
          </Button>
        </center>
        </div>

        <div className={classes.root}>
          <form noValidate autoComplete="off">

            <div className={classes.paper}><strong>ชื่อหนังสือ</strong></div>
            <TextField className={classes.textField}
              style={{ width: 400, marginLeft: 20, marginRight: -10 }}
              id="title"
              label=""
              variant="standard"
              color="secondary"
              type="string"
              size="medium"
              value={title}
              onChange={handleTitleChange}
            />

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