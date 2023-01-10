import React from "react"
import { NavLink, useNavigate } from "react-router-dom"
import LogoutIcon from "../icons/LogoutIcon"
import PencilIcon from "../icons/PencilIcon"
import MagnifierIcon from "../icons/MagnifierIcon"
import { Logout } from "../../wailsjs/go/main/App"

const Nav = () => {
  const navigate = useNavigate()

  const handleLogout = async () => {
    // TODO: make a nicer confirmation window that matches the rest of the UI.
    const confirmation = confirm("Are you sure?\n\nThis will delete the mnemonic from your system, make sure you have it written down in a secure place.")
    if (confirmation) {
      try {
        await Logout()
        navigate("/")
      } catch (err) {
        console.log(err)
      }
    }
  }

  return (
    <nav>
      <ul>
        <li>
          <NavLink to="sign" className={({ isActive }) => isActive ? 'active' : undefined }>
            <PencilIcon /> Sign
          </NavLink>
        </li>
        <li>
          <NavLink to="verify" className={({ isActive }) => isActive ? 'active' : undefined }>
            <MagnifierIcon /> Verify
          </NavLink>
        </li>
        <li title="Log Out" onClick={handleLogout} style={{ cursor: 'pointer' }}>
          <LogoutIcon /> Logout
        </li>
      </ul>
    </nav>
  )
}

export default Nav
