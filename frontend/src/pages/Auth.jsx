import React, { useState } from "react"
import { useNavigate } from "react-router-dom"
import logo from "../assets/images/logo-universal.png"
import { Login } from "../../wailsjs/go/main/App"

const Auth = () => {
  const navigate = useNavigate()
  const [error, setError] = useState("")

  const handleSubmit = async (ev) => {
    ev.preventDefault()
    setError("")
    try {
      const mnemonic = ev.target.mnemonic.value
      await Login(mnemonic)
      navigate("/dashboard")
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
