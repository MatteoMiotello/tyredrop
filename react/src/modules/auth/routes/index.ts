import { RouteObject} from "react-router-dom";
import AuthTemplate from "../AuthTemplate";
import LoginPage from "../LoginPage";
import RegisterPage from "../RegisterPage";
import ResetPasswordPage from "../ResetPasswordPage";
import ChangePasswordPage from "../ChangePasswordPage";

export const authRoutes: RouteObject = {
    path: '/auth',
    Component: AuthTemplate,
    children: [
        {
            path: 'login',
            Component: LoginPage,
        },
        {
            path: 'register',
            Component: RegisterPage,
        },
        {
            path: 'reset_password',
            Component: ResetPasswordPage
        },
        {
            path: 'change_password/:token',
            Component: ChangePasswordPage
        }
    ]
};