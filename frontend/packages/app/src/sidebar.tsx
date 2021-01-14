import React, { useState, useEffect } from 'react';
import HomeIcon from '@material-ui/icons/Home';
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
    
    
    {/* End global nav */}
    <SidebarDivider />
    <SidebarSpace />
    <SidebarDivider />
    
    {/* <SidebarUserSettings  /> */}
    <SidebarPinButton />
  </Sidebar>
  )
};
