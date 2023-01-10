import React, { useEffect, useRef } from "react"
import { useFetcher } from "react-router-dom"
import toast from "react-simple-toasts"

const ImportKey = () => {
  const fetcher = useFetcher()
  const stateRef = useRef(false)
  const formRef = useRef(null)

  useEffect(() => {
    if (fetcher.state === "submitting") {
      stateRef.current = true
    }
    if (fetcher.state === "idle" && stateRef.current) {
      toast("Key imported successfully!") 
      formRef.current?.reset()
    }
  }, [fetcher.state])

  return (
    <fetcher.Form className="import-key" method="post" ref={formRef}>
      <textarea className="black" name="newKey" cols="18" rows="4" placeholder="Paste Public Key" />
      <input className="black" type="text" name="newName" placeholder="Name" />
      <button className="outline" type="submit">Import Key</button> 
    </fetcher.Form>
  )
}

export default ImportKey
