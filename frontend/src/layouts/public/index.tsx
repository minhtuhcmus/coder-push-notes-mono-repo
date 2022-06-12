import {Outlet} from "react-router-dom";

function PublicLayout() {
  return (
    <div>
      PublicLayout
      <Outlet/>
    </div>
  )
}

export default PublicLayout;