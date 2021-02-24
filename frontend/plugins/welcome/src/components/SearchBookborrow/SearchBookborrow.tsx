import React, { FC, useState } from 'react';

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
  Button,
  Link,
  FormControl,
  TextField,
} from '@material-ui/core';
import { Link as RouterLink } from 'react-router-dom';
import { DefaultApi } from '../../api/apis';
import { EntBookborrow } from '../../api';
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

  const [name, setName] = useState(String);

  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);

  const [bookborrows, setBookborrows] = useState<EntBookborrow[]>([]);


  const handleChange = (
    event: React.ChangeEvent<{ value: any }>,
  ) => {
    const { value } = event.target;
    setName(value);
    setStatus(false);
    setBookborrows([]);
  };

  var lenBookborrow: number
  const getCheckinBookborrow = async () => {
    const res = await api.getBookborrow({ id: name })
    setBookborrows(res)
    console.log(res)
    lenBookborrow = res.length
    if (lenBookborrow > 0) {
      setStatus(true);
      setAlert(true);
    }
    else {
      setStatus(true);
      setAlert(false);
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
                <Alert severity="success" style={{ width: 400 }} onClose={() => { setStatus(false) }}>
                  <div> ค้นหาข้อมูลพบ</div>
                </Alert>

              ) : (
                  <Alert severity="error" style={{ width: 400 }} onClose={() => { setStatus(false) }}>
                    <div>
                      ไม่พบข้อมูลที่ค้นหา
                </div>
                  </Alert>

                )}
            </div>) : null}
        </ContentHeader>
        <Grid container alignItems="center" justify="center" item xs={12}>
        <Grid item xs={3}>
          <FormControl variant="outlined" className={classes.formControl}>
            <TextField

              id="User"
              label="ชื่อสมาชิกห้องสมุด"
              margin="normal"
              variant="outlined"
              type="string"
              onChange={handleChange}

            />
          </FormControl>
          </Grid>
          <Grid item xs={3}>
            <Button

              variant="contained"
              color="primary"
              size="large"
              className={classes.button}
              onClick={() => {
                getCheckinBookborrow();
              }}

              startIcon={<SearchIcon
              />}
            >
              Search
              </Button>
          </Grid>
          <Grid item xs={6}></Grid>
        </Grid>
        <Grid container spacing={1}>

        </Grid>


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
              {bookborrows.map((item: any) => (

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