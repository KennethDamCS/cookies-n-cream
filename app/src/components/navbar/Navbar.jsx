import React from 'react'
import { styled } from '@mui/material/styles';
import HomeRoundedIcon from '@mui/icons-material/HomeRounded';
import ExploreOutlinedIcon from '@mui/icons-material/ExploreOutlined';
import InboxRoundedIcon from '@mui/icons-material/InboxRounded';
import SettingsIcon from '@mui/icons-material/Settings';
import { Button } from '@mui/material';
import IconButton from '@mui/material/IconButton';
import Stack from '@mui/material/Stack';

import './Navbar.css'

function Navbar () {

  return (
   <section className="navbar">
      <Stack direction="row" spacing={1}>
        <IconButton aria-label="dashboard" color="inherit">
          <HomeRoundedIcon />
        </IconButton>
        <IconButton aria-label="dashboard" color="inherit">
          <ExploreOutlinedIcon />
        </IconButton>
        <IconButton aria-label="dashboard" color="inherit">
          <InboxRoundedIcon />
        </IconButton>
        <IconButton aria-label="dashboard" color="inherit">
          <SettingsIcon />
        </IconButton>
        <Button variant="outlined" size="small" color="inherit">
          Login
        </Button>
      </Stack>

  </section>
  )

}

export default Navbar;