import {createContext} from "react";
import {SignInRequest, SignUpRequest} from "../interfaces";

interface AuthContextType {
  isLogin: boolean;
  signup: (request: SignUpRequest) => any;
  signin: (request: SignInRequest) => any;
  signout: () => void;
}

let AuthContext = createContext<AuthContextType>(null!);

export default AuthContext;