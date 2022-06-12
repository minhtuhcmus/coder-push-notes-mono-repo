import {useState, useEffect} from "react";
import * as Auth from "../../requests/auth";
import AuthContext from "../../contexts/auth_context";
import {SignInRequest, SignUpRequest} from "../../interfaces";

function AuthProvider({ children }: { children: React.ReactNode }) {
  let [isLogin, setIsLogin] = useState<boolean>(false);

  useEffect(() => {
    async function authentication(): Promise<boolean> {
      return Auth.VerifyToken();
    }
    authentication().then(res => setIsLogin(res));
  }, []);

  let signin = async (request: SignInRequest) => {
    let res = await Auth.SignIn(request);
    console.log("res", res)
    if (res.data == null) {
      setIsLogin(false);
    }
    setIsLogin(true);
  };

  let signup = async (request: SignUpRequest) => {
    let res = await Auth.SignUp(request);
    console.log("res", res)
    if (res.data == null) {
      setIsLogin(false);
    }
    setIsLogin(true);
  }

  let signout = () => {
    setIsLogin(false);
  };

  let value = { isLogin, signup, signin, signout };

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
}

export {
  AuthProvider
};