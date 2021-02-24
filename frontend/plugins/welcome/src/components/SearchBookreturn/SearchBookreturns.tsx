import React, { FC } from 'react';

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
    TextField,
    Grid,
    FormControl,
    Button,
} from '@material-ui/core';
import { DefaultApi } from '../../api/apis';
import { EntBookreturn } from '../../api/models/EntBookreturn';
import Snackbar from '@material-ui/core/Snackbar';
import Alert from '@material-ui/lab/Alert';



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
    const [name, setName] = React.useState(String)
    
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
   
    
    const [bookreturn, setBookreturn] = React.useState<EntBookreturn[]>([])
    const getBookreturns = async () => {
        const res = await api.getBookreturn({id:name})
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

    // set data to object and setName
    const handleChange = (
        event: React.ChangeEvent<{ value: any }>,
    ) => {
        const { value } = event.target;
        setName(value);
    };


    // function seach data
    function search() {
        getBookreturns();
    }

   

    return (
        <Page theme={pageTheme.home}>
            <Header style={HeaderCustom} title={`ระบบค้นหารายการคืนหนังสือ`}>
                
            </Header>
            <Content>
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
                    <FormControl variant="outlined" >

                        <Button
                            variant="contained"
                            color="secondary"
                            size="large"
                            onClick={search}
                        >
                            search
                        </Button>

                        </FormControl>
                    </Grid>
                    <Grid item xs={6}></Grid>
                </Grid>
                <Grid container spacing={1}>
                    
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
                            {bookreturn.map(item => (
                               
                                <TableRow key={item.id}>
                                    
                                    <TableCell align="center">{item.id}</TableCell>
                                    <TableCell align="center">{item.edges?.user?.uSERNAME}</TableCell>
                                    <TableCell align="center">{item.edges?.mustreturn?.id}</TableCell>                         
                                    <TableCell align="center">{item.edges?.mustreturn?.edges?.book?.bookName}</TableCell>
                                    <TableCell align="center">{item.edges?.location?.lOCATIONNAME}</TableCell>
                                    <TableCell align="center">{item.dAMAGEDPOINT=== undefined? "0":item.dAMAGEDPOINT 
                                    }</TableCell>
                                    <TableCell align="center">{item.dAMAGEDPOINTNAME}</TableCell>
                                    <TableCell align="center">{item.lOST}</TableCell>
                                    <TableCell align="center">{item.rETURNTIME}</TableCell>
                                    
                                </TableRow>
                            ))}
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