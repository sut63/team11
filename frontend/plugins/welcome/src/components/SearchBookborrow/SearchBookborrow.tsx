import React, { FC, useEffect, useState } from 'react';

import { makeStyles } from '@material-ui/core/styles';
import { Content, ContentHeader, Header, Page, pageTheme } from '@backstage/core';
import Typography from '@material-ui/core/Typography';
import Divider from '@material-ui/core/Divider';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow'
import Paper from '@material-ui/core/Paper';
import { Alert } from '@material-ui/lab';
import {
  Grid,
  Select,
  MenuItem,
  Button,
  Link,
} from '@material-ui/core';
import { Link as RouterLink } from 'react-router-dom';
import { DefaultApi } from '../../api/apis';
import { EntBookborrow, EntUser } from '../../api';
import { SearchIcon } from '@material-ui/data-grid';


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
  bottom: {
    marginLeft: theme.spacing(3),
    marginTop: theme.spacing(1),
    marginBottom: theme.spacing(2),
  },
  divider: {
    margin: theme.spacing(2, 0),
  },
  table: {
    minWidth: 650,
  },
  button: {
    margin: theme.spacing(1),
  },
}));


const SearchBookborrows: FC<{}> = () => {
  const classes = useStyles();
  const api = new DefaultApi();

  const [users, setUsers] = useState<EntUser[]>([]);
  const [loading, setLoading] = useState(true);
  const [userid, setUserid] = useState(Number);

  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);

  const [bookborrows, setBookborrows] = useState<EntBookborrow[]>([]);

  useEffect(() => {
    const getUsers = async () => {
      const pr = await api.listUser();
      setLoading(false);
      setUsers(pr);
    };
    getUsers();
  }, [loading]);

  const UserID_handleChange = (event: any) => {
    setUserid(event.target.value);
    setStatus(false);
    setBookborrows([]);
  }

  var lenBookborrow: number
  const getCheckinBookborrw = async () => {
    const res = await api.listBookborrow({limit:10 , offset:0})
    setBookborrows(res)
    for(let i = 0; i < res.length ; i++){
      if(res[i].edges?.user?.id == userid){
        lenBookborrow = 1
        break
      }
      else{
        lenBookborrow = 0
      }
    }
    console.log(res)
    console.log(userid)
    if (lenBookborrow == 0) {
      setStatus(true);
      setAlert(false);
    }
    else {
      setStatus(true);
      setAlert(true);
    }

  }

  return (
    <Page theme={pageTheme.home}>
      <Header title={`ระบบค้นหารายการยืมหนังสือ`}>
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
      </Header>
      <Content>
        <ContentHeader title={'ชื่อผู้ยืมหนังสือที่ต้องการค้นหา'}>
          {status ? (
            <div>
              {alert ? (
                <Alert severity="success" style={{ width: 400 }} onClose={() => { setStatus(false)}}>
                 <div> ค้นหาข้อมูลพบ</div>
                </Alert>
                
          ) : (
                <Alert severity="error" style={{ width: 400 }} onClose={() => { setStatus(false)}}>
                <div>
                  ไม่พบข้อมูลที่ค้นหา
                </div>
                </Alert>
                
          )}
           </div>) : null}
        </ContentHeader>
        <Grid item xs={2}></Grid>
        <Grid item xs={2}></Grid>
        <Grid item xs={2}>
          <Select
            labelId="user_id-label"
            label="User"
            id="UserID"
            onChange={UserID_handleChange}
            style={{ width: 300 }}
            value={userid}
            variant="outlined"
          >
            {users.map((item: EntUser) =>
              <MenuItem key={item.id} value={item.id}>{item.uSERNAME}</MenuItem>)}
          </Select>
        </Grid>
        <Grid item xs={2}></Grid>
        <Grid item xs={2}> </Grid>

        <Grid item xs={2}></Grid>
        <Grid item xs={2}> </Grid>
        <Grid item xs={2}></Grid>
        <Grid item xs={2}>
          <Button

            variant="contained"
            color="primary"
            size="large"
            className={classes.button}
            onClick={() => {
              getCheckinBookborrw();
            }}

            startIcon={<SearchIcon
            />}
          >
            Search
              </Button>
        </Grid>
        <Grid item xs={2}></Grid>
        <Grid item xs={2}> </Grid>
        <Grid item xs={12}></Grid>
        <Grid item xs={12}></Grid>


        <Grid item xs={12}>
          <Divider className={classes.divider} />
          <Typography variant="subtitle1" gutterBottom>
            ตารางข้อมูลการยืมหนังสือ:
                        </Typography>
        </Grid>
        <TableContainer component={Paper}>
          <Table className={classes.table} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell align="center">id</TableCell>
                <TableCell align="center">ชื่อผู้ยืมหนังสือ</TableCell>
                <TableCell align="center">ชื่อหนังสือที่ยืม</TableCell>
                <TableCell align="center">จุดบริการที่ยืมหนังสือ</TableCell>
                <TableCell align="center">วันที่ยืม</TableCell>
                <TableCell align="center">จำนวนวันที่ยืม</TableCell>
                <TableCell align="center">วิธีการรับหนังสือ</TableCell>
                <TableCell align="center">เบอร์โทรศัพท์</TableCell>
                <TableCell align="center">สถานะการยืม</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {bookborrows.filter((filter: any) => filter.edges.user?.id == userid).map((item: any) => (

                <TableRow key={item.id}>
                  <TableCell align="center">{item.id}</TableCell>
                  <TableCell align="center">{item.edges?.user?.uSERNAME}</TableCell>
                  <TableCell align="center">{item.edges?.book?.bookName}</TableCell>
                  <TableCell align="center">{item.edges?.servicepoint?.cOUNTERNUMBER}</TableCell>
                  <TableCell align="center">{item.bORROWDATE}</TableCell>
                  <TableCell align="center">{item.dAYOFBORROW}</TableCell>
                  <TableCell align="center">{item.pICKUP}</TableCell>
                  <TableCell align="center">{item.pHONENUMBER}</TableCell>
                  <TableCell align="center">{item.edges?.status?.sTATUSNAME}</TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Content>
    </Page>
  );
};

export default SearchBookborrows;