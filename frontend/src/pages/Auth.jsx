import React, { useEffect } from "react"
import { Form, redirect, useNavigate } from "react-router-dom"
import logo from "../assets/images/logo-universal.png"
import { Login, IsLoggedIn } from "../../wailsjs/go/main/App"

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
      <Form method="post">
        <input
          autoFocus
          type="text"
          name="mnemonic"
          autoComplete="off"
          placeholder="12-word mnemonic"
          className="wide"
        />
      </Form>
    </main>
  )
}

export default Auth

export async function action({ request }) {
  const formData = await request.formData()
  await Login(formData.get("mnemonic"))
  return redirect("/dashboard/sign")
}
