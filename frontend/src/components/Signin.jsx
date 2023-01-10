import React from "react"
import { Form, redirect, Link } from "react-router-dom"
import { Login } from "../../wailsjs/go/main/App"

const Signin = () => {
  return (
    <div>
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
      <Link to="/">&larr; Go Back</Link>
    </div>
  )
}

export default Signin

export async function action({ request }) {
  const formData = await request.formData()
  await Login(formData.get("mnemonic"))
  return redirect("/dashboard/sign")
}
