import { RouteObject} from "react-router-dom";
import AuthTemplate from "../AuthTemplate";
import LoginPage from "../LoginPage";
import RegisterPage from "../RegisterPage";

export const loginRoute: RouteObject = {
    path: 'login',
    Component: LoginPage,
};

export const registerRoute: RouteObject = {
    path: 'register',
    Component: RegisterPage,
};

export const authRoutes: RouteObject = {
    path: '/auth',
    Component: AuthTemplate,
    children: [
        loginRoute,
        registerRoute
    ]
};