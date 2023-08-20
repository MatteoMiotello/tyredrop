import React, {HTMLAttributes, PropsWithChildren} from "react";
import {twMerge} from "tailwind-merge";

type StatProps = HTMLAttributes<HTMLDivElement> & PropsWithChildren
const Stat: React.FC<StatProps> = ({children, className, ...props}) => {
    const classNames = twMerge(
        'stat shadow bg-base-100',
        className
    );

    return <div className={classNames} {...props}>
        {children}
    </div>;
};

type FigureProps = {
    className?: string
} & PropsWithChildren
const Figure: React.FC<FigureProps> = ({children, className}) => {
    const classNames = twMerge(
        'stat-figure text-4xl',
        className
    );

    return <div className={classNames}>
        {children}
    </div>;
};

type TitleProps = {
    className?: string
} & PropsWithChildren
const Title: React.FC<TitleProps> = ({children, className}) => {
    const classNames = twMerge(
        'stat-title',
        className
    );

    return <div className={classNames}>
        {children}
    </div>;
};

type ValueProps = {
    className?: string
} & PropsWithChildren
const Value: React.FC<ValueProps> = ({children, className}) => {
    const classNames = twMerge(
        'stat-value',
        className
    );

    return <div className={classNames}>
        {children}
    </div>;
};

type DescProps = {
    className?: string
} & PropsWithChildren
const Desc: React.FC<DescProps> = ({children, className}) => {
    const classNames = twMerge(
        'stat-desc',
        className
    );

    return <div className={classNames}>
        {children}
    </div>;
};

export default Object.assign(Stat, {
    Figure,
    Title,
    Value,
    Desc
});