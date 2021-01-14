import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntUser } from '../../api/models/EntUser'; // import interface User
import { EntRoominfo } from '../../api/models/EntRoominfo'; // import interface Roominfo
const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function ComponentsTable() {
 const classes = useStyles();
 const api = new DefaultApi();
 
 const [roominfos, setRoominfos] = React.useState<EntRoominfo[]>([]);
 const [loading, setLoading] = useState(true);
 
 useEffect(() => {
    
    
 
    const getRoominfos = async () => {
    const res = await api.listRoominfo({ limit: 10, offset: 0 });
    setRoominfos(res);
  };
  getRoominfos();
}, [loading]);
 
 
 return ( 
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
           <TableCell align="right">เลขห้อง</TableCell>
           <TableCell align="right">ประเภทห้อง</TableCell>
           <TableCell align="right">เวลา</TableCell>
           <TableCell align="right">สถานะ</TableCell>
         </TableRow>
       </TableHead>
       <TableBody>
         {roominfos.map(item => (
           <TableRow key={item.id}>
             <TableCell align="right">{"ห้อง"+item.roomNo}</TableCell>
             <TableCell align="right">{item.roomType}</TableCell>
             <TableCell align="right">{item.roomTime}</TableCell>
             <TableCell align="right">{item.roomStatus}</TableCell>
             <TableCell align="right">
               
             </TableCell>
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
 );
}
