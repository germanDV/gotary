import React, { useState } from "react"
import { Form, redirect } from "react-router-dom"
import { SelectFile, VerifySignature } from "../../wailsjs/go/main/App"

const Verifier = () => {
  const [filePath, setFilePath] = useState("")
  const [signaturePath, setSignaturePath] = useState("")
  const [err, setErr] = useState("")

  const selectFile = async () => {
    try {
      const p = await SelectFile() 
      setFilePath(p)
    } catch (err) {
      setErr(err)
    }
  }

  const selectSignature = async () => {
    try {
      const p = await SelectFile() 
      setSignaturePath(p)
    } catch (err) {
      setErr(err)
    }
  }

  return (
    <div>
      <Form method="post">
        <section>
          <h2>Select File To Verify</h2>
          <div className="file-chooser">
            <button type="button" className="outline" onClick={selectFile}>Open</button>
            <div>{filePath}</div>
          </div>
          <input type="hidden" name="filePath" value={filePath} />
        </section>
        <section>
          <h2>Select Signature</h2>
          <div className="file-chooser">
            <button type="button" className="outline" onClick={selectSignature}>Open</button>
            <div>{signaturePath}</div>
          </div>
          <input type="hidden" name="signaturePath" value={signaturePath} />
        </section>
        <section>
          <h2>Public Key Of Signer</h2>
          <textarea name="key" className="full" rows="3" placeholder="Paste one of your contacts' key" />
        </section>
        <button type="submit" className="fill">
          Verify signature
        </button>
      </Form>
      {err ? <p className="error">{err}</p> : null}
    </div>
  )
}

export default Verifier

export async function action({ request }) {
    const formData = await request.formData()
    const filePath = formData.get("filePath")
    const signaturePath = formData.get("signaturePath")
    const key = formData.get("key")
    await VerifySignature(filePath, signaturePath, key)
    return redirect(`/success?file=${encodeURI(filePath)}`)
}
