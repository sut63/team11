import React from 'react';
import HomeIcon from '@material-ui/icons/Home';
import CreateComponentIcon from '@material-ui/icons/AddCircleOutline';
import YouTube from '@material-ui/icons/YouTube';
import SignOut from '@material-ui/icons/Settings';
import Book from '@material-ui/icons/MenuBook';
import VideoOnDemand from '@material-ui/icons/DesktopWindows';
import Room from '@material-ui/icons/EventSeat';
import Research from '@material-ui/icons/NoteAdd';
import Borrow from '@material-ui/icons/AddShoppingCart';
import Return from '@material-ui/icons/AssignmentReturn';
import {
  Sidebar,
  SidebarItem,
  SidebarDivider,
  SidebarSpace,
  SidebarUserSettings,
  SidebarThemeToggle,
  SidebarPinButton,
} from '@backstage/core';

export const AppSidebar = () => (
  <Sidebar>
    <SidebarDivider />
    {/* Global nav, not org-specific */}
    <SidebarItem icon={HomeIcon} to="" text="Home" />
    {/* <SidebarItem icon={CreateComponentIcon} to="create" text="Create..." />
    <SidebarItem icon={CreateComponentIcon} to="welcome" text="Welcome" /> */}
    <SidebarItem
      icon={YouTube}
      to="playlist_video"
      text="Playlist Video"
    />
    <SidebarItem
      icon={CreateComponentIcon}
      to="watch_video"
      text="Watch Video"
    />

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
      to="Borrow"
      text="ยืมหนังสือ"
    />
    <SidebarItem
      icon={Return}
      to="Return"
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
);
