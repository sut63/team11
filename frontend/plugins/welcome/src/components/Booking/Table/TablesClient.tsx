import { DefaultApi } from '../../../api/apis';
import { EntClientEntity } from '../../../api/models/EntClientEntity';

import React, { useState, useEffect } from 'react';
import { withStyles, Theme, createStyles, makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';

const StyledTableCell = withStyles((theme: Theme) =>
  createStyles({
    head: {
      backgroundColor: theme.palette.common.black,
      color: theme.palette.common.white,
    },
    body: {
      fontSize: 14,
    },
  }),
)(TableCell);

const StyledTableRow = withStyles((theme: Theme) =>
  createStyles({
    root: {
      '&:nth-of-type(odd)': {
        backgroundColor: theme.palette.action.hover,
      },
    },
  }),
)(TableRow);

const useStyles = makeStyles({
  table: {
    MaxWidth: 200,
  },
});

export default function CustomizedTables() {
  const classes = useStyles();
  const api = new DefaultApi();

    const [clients, setClients] = useState<EntClientEntity[]>(Array);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const getUsers = async () => {
            const res = await api.listCliententity();
            setLoading(false);
            setClients(res);
        };
        getUsers();
    }, [loading]);

  return (
    <TableContainer component={Paper}>
      <Table className={classes.table} aria-label="customized table">
          
        <TableHead>
          <TableRow>
            <StyledTableCell>ลำดับ (No.)</StyledTableCell>
            <StyledTableCell align="left">เครื่อง (Video on demand Client)</StyledTableCell>
            <StyledTableCell align="left">สถานะ</StyledTableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {clients.map((row:EntClientEntity) => (
            <StyledTableRow key={row.id}>
              <StyledTableCell component="th" scope="row">
                {row.id}
              </StyledTableCell>
              <StyledTableCell align="left">{row.cLIENTNAME}</StyledTableCell>
              <StyledTableCell align="left">{row.edges?.state?.sTATUSNAME}</StyledTableCell>
            </StyledTableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
}
