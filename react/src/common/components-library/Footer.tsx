import React, {HTMLAttributes, PropsWithChildren} from "react";
import {twMerge} from "tailwind-merge";

export type FooterColumn = {
    title?: string
    key: number | string
    links: FooterLink[]
}

export type FooterLink = {
    title: string
    key: number | string
    url?: string
}

type FooterProps = PropsWithChildren

const Footer: React.FC<FooterProps> = ({children} ) => {
    return <footer className="footer p-10 bg-neutral text-neutral-content w-full">
        {children}
    </footer>;
};

type LinkProps = HTMLAttributes<HTMLSpanElement> & PropsWithChildren

const Element: React.FC<LinkProps> = ({className, children, ...props} ) => {
    const classNames = twMerge( 
        'footer-title',
        className
    );
    
    return <span className={classNames} {...props}>
        {children}
    </span>;
};

type ColumnProps = {
    className?: string
} & PropsWithChildren
const Column: React.FC<ColumnProps> = ({ className, children}) => {
    return <div className={className}>
        {children}
    </div>;
};

export default Object.assign( Footer, {
    Element,
    Column
});