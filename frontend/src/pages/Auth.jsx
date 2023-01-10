import React, { useEffect } from "react"
import { useNavigate, Outlet, Link } from "react-router-dom"
import logo from "../assets/images/logo-universal.png"
import { IsLoggedIn } from "../../wailsjs/go/main/App"

const Auth = () => {
  const navigate = useNavigate()

  useEffect(() => {
    async function checkUserStatus() {
      const isLoggedIn = await IsLoggedIn()
      if (isLoggedIn) {
        navigate("/dashboard/sign")
      }
    }
    checkUserStatus()
  }, [])

  return (
    <main>
      <img src={logo} id="logo" alt="logo"/>
      <h1>Gotary</h1>
      <Outlet />
    </main>
  )
}

export default Auth

