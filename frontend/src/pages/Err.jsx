import React from "react"
import { useRouteError, Link } from "react-router-dom"
import ShieldExclamationIcon from "../icons/ShieldExclamationIcon"

const Err = () => {
  const err = useRouteError()

  return (
    <main>
      <ShieldExclamationIcon />
      <h1>ERROR!</h1>
      <h4 className="error">{JSON.stringify(err)}</h4>
      <Link to="/">&larr; Go Back</Link>
    </main>
  )
}

export default Err
