import React, { useState, useEffect } from 'react';
import HomeIcon from '@material-ui/icons/Home';
import SignOut from '@material-ui/icons/Settings';
import Book from '@material-ui/icons/MenuBook';
import VideoOnDemand from '@material-ui/icons/DesktopWindows';
import Room from '@material-ui/icons/EventSeat';
import Research from '@material-ui/icons/NoteAdd';
import Borrow from '@material-ui/icons/AddShoppingCart';
import Return from '@material-ui/icons/AssignmentReturn';
import { Link as RouterLink } from 'react-router-dom';
import {
  Sidebar,
  SidebarItem,
  SidebarDivider,
  SidebarSpace,
  SidebarPinButton,
} from '@backstage/core';


export const AppSidebar =  () => {
  const [loading, setLoading] = useState(true);
  useEffect(() => {

  }, [loading]);
  return(
  <Sidebar>
    <SidebarDivider />
    {/* Global nav, not org-specific */}
    <SidebarItem icon={HomeIcon} to="" text="Home" />
    {}
    <SidebarItem
      icon={Book}
      to="Book"
      text="เพิ่มหนังสือ"
    />
    <SidebarItem
      icon={VideoOnDemand}
      to="VideoOnDemand"
      text="จองเครื่องรับชม VideoOnDemand"
    />
    <SidebarItem
      icon={Room}
      to="Room"
      text="จองห้องติว"
    />
    <SidebarItem
      icon={Research}
      to="Research"
      text="เพิ่มงานวิจัย"
    />
    <SidebarItem
      icon={Borrow}
      to="Bookborrow"
      text="ยืมหนังสือ"
    />
    <SidebarItem
      icon={Return}
      to="Bookreturn"
      text="คืนหนังสือ"
    />
    
    
    {/* End global nav */}
    <SidebarDivider />
    <SidebarSpace />
    <SidebarDivider />
    <SidebarItem
      icon={SignOut}
      to="sign_out"
      text="Sign Out"
    />
    {/* <SidebarUserSettings  /> */}
    <SidebarPinButton />
  </Sidebar>
  )
};
