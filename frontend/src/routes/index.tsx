import {AuthProvider} from "../components/commons/auth_provider";
import {Routes, Route} from "react-router-dom";
import PublicLayout from "../layouts/public";
import PrivateLayout from "../layouts/private";

function CustomRoutes() {
  return (
    <AuthProvider>
      <Routes>
        <Route path={"/"} element={<PublicLayout/>}>
          <Route path='*' element={<div>Todo Not Found Page</div>}/>
          <Route path={"signin"} element={<div>Todo Sign In Page</div>}/>
          <Route path={"signup"} element={<div>Todo Sign Up Page</div>}/>
          <Route path={"/"} element={<div>Todo Home Page</div>}/>
        </Route>

        <Route path={"/private"} element={<PrivateLayout/>}>
          <Route path={"note-app"} element={<div>Todo Note App</div>}/>
          <Route path={"profile"} element={<div>Todo Profile Page</div>}/>
        </Route>
      </Routes>
    </AuthProvider>
  )
}

export default CustomRoutes;