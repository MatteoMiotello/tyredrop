import clsx from "clsx";
import React, {HTMLAttributes, PropsWithChildren, ReactNode} from "react";
import {twMerge} from "tailwind-merge";

type MenuProps = {
    direction?: 'horizontal' | 'vertical' | undefined
} & HTMLAttributes<HTMLUListElement> & PropsWithChildren
const Menu: React.FC<MenuProps> = ({className, children, direction, ...props}) => {
    if (!direction) {
        direction = 'vertical';
    }

    const classNames = twMerge(
        'menu',
        className,
        clsx({
            'menu-horizontal': direction == 'horizontal',
        })
    );

    return <ul {...props} className={classNames}>
        {children}
    </ul>;
};

type MenuItemProps = HTMLAttributes<HTMLLIElement> & PropsWithChildren
const Item: React.FC<MenuItemProps> = ({children, ...props}) => {
    return <li {...props}>
        {children}
    </li>;
};

type SubmenuProps = {
    title: string
} & PropsWithChildren & HTMLAttributes<HTMLUListElement>
const Submenu: React.FC<SubmenuProps> = ({title, children, ...props}) => {
    return <details open>
        <summary>
            {title}
        </summary>
        <Menu {...props}>
            {children}
        </Menu>
    </details>;
};


type DropdownProps = {
    label: string | ReactNode
} & PropsWithChildren
const Dropdown: React.FC<DropdownProps> = ({label, children}) => {
    return <details className="dropdown">
        <summary tabIndex={0} className="btn">{label}</summary>
        <Menu tabIndex={0} className="dropdown-content p-2 shadow rounded-box w-52 bg-base-100">
            {children}
        </Menu>
    </details>;
};

export default Object.assign(Menu, {
    Item: Item,
    Submenu: Submenu,
    Dropdown: Dropdown
});