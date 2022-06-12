import {useContext} from "react";
import AuthContext from "../contexts/auth_context";

function useAuth() {
  return useContext(AuthContext);
}

export default useAuth;

