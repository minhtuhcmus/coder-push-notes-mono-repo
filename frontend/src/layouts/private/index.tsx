import {Outlet} from "react-router-dom";
import RequireAuth from "../../components/commons/require_auth";

function PrivateLayout() {
  return (
    <RequireAuth>
      <>
        PrivateLayout
        <Outlet/>
      </>
    </RequireAuth>
  )
}

export default PrivateLayout;