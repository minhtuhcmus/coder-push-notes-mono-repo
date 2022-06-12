import {SignInRequest, SignUpRequest} from "../interfaces";

async function SignIn(request: SignInRequest) : Promise<any> {
  //Todo: To check if user is in db and assign an access token (jwt) in cookie
  // (fyi: this access token will contain userID to verify each request)
}

async function SignUp(request: SignUpRequest) : Promise<any> {
  //Todo: To SignUp new user
}

async function VerifyToken() : Promise<boolean> {
  //Todo: To verify access token in cookie
  return false
}

export {
  SignIn,
  SignUp,
  VerifyToken
}