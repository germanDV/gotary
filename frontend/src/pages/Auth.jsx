import React, { useState, useEffect } from "react"
import { useNavigate } from "react-router-dom"
import logo from "../assets/images/logo-universal.png"
import { Login, IsLoggedIn } from "../../wailsjs/go/main/App"

const Auth = () => {
  const navigate = useNavigate()
  const [error, setError] = useState("")

  useEffect(() => {
    async function checkUserStatus() {
      const isLoggedIn = await IsLoggedIn()
      if (isLoggedIn) {
        navigate("/dashboard/sign")
      }
    }
    checkUserStatus()
  }, [])

  const handleSubmit = async (ev) => {
    ev.preventDefault()
    setError("")
    try {
      const mnemonic = ev.target.mnemonic.value
      await Login(mnemonic)
      navigate("/dashboard/sign")
    } catch (err) {
      setError(typeof err === "string" ? err : JSON.stringify(err))
    }
  }

  return (
    <main>
      <img src={logo} id="logo" alt="logo"/>
      <h1>Auth Page</h1>
      <form onSubmit={handleSubmit}>
        <input
          autoFocus
          type="text"
          name="mnemonic"
          autoComplete="off"
          placeholder="12-word mnemonic"
          className="wide"
        />
      </form>
      <div className="error">{error}</div>
    </main>
  )
}

export default Auth
