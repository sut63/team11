import React, { FC, useEffect } from 'react';

import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme } from '@backstage/core';
import Typography from '@material-ui/core/Typography';
import Divider from '@material-ui/core/Divider';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow'
import Paper from '@material-ui/core/Paper';

import {
    Grid,
    FormControl,
    Select,
    InputLabel,
    MenuItem,
    Avatar,
    Button,
} from '@material-ui/core';
import { DefaultApi } from '../../api/apis';
import { EntBookreturn } from '../../api/models/EntBookreturn';
import Snackbar from '@material-ui/core/Snackbar';
import Alert from '@material-ui/lab/Alert';
import { EntBookborrow, EntUser } from '../../api';


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
}));


const SearchBookreturns: FC<{}> = () => {
    const classes = useStyles();
    const api = new DefaultApi();

    

    // User
    const [idUser, setIduser] = React.useState<number>(0)
    const [user, setUser] = React.useState<EntUser[]>([])
    const getUser = async () => {
        const res = await api.listUser()
        setUser(res)
    }

    // alert setting
    const [open, setOpen] = React.useState(false);
    const [fail, setFail] = React.useState(false);

    //close alert 
    const handleClose = (event?: React.SyntheticEvent, reason?: string) => {
        if (reason === 'clickaway') {
            return;
        }
        setFail(false);
        setOpen(false);
    };


    // Bookreturn
    var lenreturn : number
    const [bookborrow, setBookBorrow] = React.useState<EntBookborrow[]>([])
    const getBookBorrows = async () => {
        const res = await api.listBookborrow({})
        setBookBorrow(res)
    }
    console.log("xx",bookborrow);
    
    const [bookreturn, setBookreturn] = React.useState<EntBookreturn[]>([])
    const getBookreturns = async () => {
        const res = await api.getBookreturn({id:idUser})
        setBookreturn(res)
        console.log(res);
        
        lenreturn = res.length
        if (lenreturn > 0){
            setOpen(true)
        }else{
            setFail(true)
        }   
    }
   console.log("bb",bookreturn);
   
    

    // Lifecycle Hooks
    useEffect(() => {
        getUser();
        getBookBorrows();
    }, []);
    


    // set data to object and setIduser
    const handleChange = (
        event: React.ChangeEvent<{ name?: string; value: any }>,
    ) => {
        const name = event.target.name as keyof typeof SearchBookreturns;
        const { value } = event.target;
        setIduser(value);
    };


    // function seach data
    function seach() {
        getBookreturns();
    }

   

    return (
        <Page theme={pageTheme.home}>
            <Header style={HeaderCustom} title={`ระบบค้นหารายการคืนหนังสือ`}>
                <Avatar alt="Remy Sharp" src="../../image/account.jpg" />
                
            </Header>
            <Content>
                <Grid container spacing={1}>
                    <Grid item xs={1}>
                        <div className={classes.paper}><h3>User</h3></div>
                    </Grid>
                    <Grid item xs={3}>
                        <FormControl variant="outlined" className={classes.formControl}>
                            <InputLabel>select customer</InputLabel>
                            <Select
                                name="User"
                                value={idUser || ''} // (undefined || '') = ''
                                onChange={handleChange}
                            >
                                {user.map(item => {
                                    return (
                                        <MenuItem key={item.id} value={item.id}>
                                            {item.uSERNAME}
                                        </MenuItem>
                                    );
                                })}
                            </Select>
                        </FormControl>
                    </Grid>
                    <Grid item xs={6}>
                        <Button
                            variant="contained"
                            color="secondary"
                            size="large"
                            onClick={seach}
                        >
                            seach
                        </Button>
                    </Grid>
                </Grid>

                <Grid item xs={12}>
                    <Divider className={classes.divider} />
                    <Typography variant="subtitle1" gutterBottom>
                        ตารางการคืนหนังสือ:
                        </Typography>
                </Grid>
                <TableContainer component={Paper}>
                    <Table className={classes.table} aria-label="simple table">
                        <TableHead>
                            <TableRow>
                                <TableCell align="center">id</TableCell>
                                <TableCell align="center">ชื่อสมาชิก</TableCell>
                                <TableCell align="center">เลขรายการยืมหนังสือ</TableCell>
                                <TableCell align="center">ชื่อหนังสือที่ยืม</TableCell>
                                <TableCell align="center">สถานที่คืนหนังสือ</TableCell>
                                <TableCell align="center">จำนวนจุดที่เสียหาย</TableCell>
                                <TableCell align="center">ชื่อตำแหน่งที่เสียหาย</TableCell>
                                <TableCell align="center">กรณีหนังสือหาย</TableCell>
                                <TableCell align="center">เวลา</TableCell>
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {bookreturn.map(item => (bookborrow.filter(bb => bb.id === item.edges?.mustreturn?.id).map(item2 => (
                               
                                <TableRow key={item2.id}>
                                    
                                    <TableCell align="center">{item.id}</TableCell>
                                    <TableCell align="center">{item.edges?.user?.uSERNAME}</TableCell>
                                    <TableCell align="center">{item.edges?.mustreturn?.id}</TableCell>                         
                                    <TableCell align="center">{item2.edges?.book?.bookName}</TableCell>
                                    <TableCell align="center">{item.edges?.location?.lOCATIONNAME}</TableCell>
                                    <TableCell align="center">{item.dAMAGEDPOINT=== undefined? "0":item.dAMAGEDPOINT 
                                    }</TableCell>
                                    <TableCell align="center">{item.dAMAGEDPOINTNAME}</TableCell>
                                    <TableCell align="center">{item.lOST}</TableCell>
                                    <TableCell align="center">{item.rETURNTIME}</TableCell>
                                    
                                </TableRow>
                            ))))}
                        </TableBody>
                    </Table>
                </TableContainer>
                <Snackbar open={open} autoHideDuration={6000} onClose={handleClose}>
            <Alert onClose={handleClose} severity="success">
              ค้นหาสำเร็จ
          </Alert>
          </Snackbar>

          <Snackbar open={fail} autoHideDuration={6000} onClose={handleClose}>
            <Alert onClose={handleClose} severity="error">
              ไม่มีรายการคืนหนังสือ
          </Alert>
          </Snackbar>
            </Content>
        </Page>
    );
};

export default SearchBookreturns;