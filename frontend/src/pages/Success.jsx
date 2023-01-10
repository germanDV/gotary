import React from "react"
import { useSearchParams, Link } from "react-router-dom"
import ShieldCheckIcon from "../icons/ShieldCheckIcon"

const Success = () => {
  const [searchParams] = useSearchParams()
  const filePath = decodeURI(searchParams.get("file") ?? "")

  return (
    <main>
      <ShieldCheckIcon />
      <h1>Success!</h1>
      <h2>
        File <span className="accent">{filePath}</span>
        <br />
        has been signed with the key you provided.
      </h2>
      <Link to="/dashboard/verify">&larr; Go Back</Link>
    </main> 
  )
}

export default Success
