import React from "react"
import { useNavigate } from "react-router-dom"
import LogoutIcon from "../icons/LogoutIcon"
import { Logout } from "../../wailsjs/go/main/App"

const Out = () => {
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
    <div className="logout" onClick={handleLogout} title="Log Out">
      <LogoutIcon />
    </div>
  )
}

export default Out
