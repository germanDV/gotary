import React from "react"
import { Link } from "react-router-dom"

const AuthLinks = () => {
  return (
    <nav className="auth-links">
      <Link to="newmnemonic">Generate Mnemonic</Link>
      <Link to="signin">I Got A Mnemonic</Link>
    </nav>
  )
}

export default AuthLinks
