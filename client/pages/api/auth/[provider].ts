import { NextApiRequest, NextApiResponse } from "next";
import firebase from "@/lib/firebase";

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse
) {
  const { provider } = req.query;
  const providerId = provider as string;

  try {
    if (providerId === "google") {
      const provider = new firebase.auth.GoogleAuthProvider();
      await firebase.auth().signInWithRedirect(provider);
    } else if (providerId === "github") {
      const provider = new firebase.auth.GithubAuthProvider();
      await firebase.auth().signInWithRedirect(provider);
    } else if (providerId === "facebook") {
      const provider = new firebase.auth.FacebookAuthProvider();
      await firebase.auth().signInWithRedirect(provider);
    } else {
      throw new Error("Invalid provider");
    }
  } catch (error) {
    console.error("Error during authentication:", error);
    res.status(500).send("Authentication failed");
  }
}
