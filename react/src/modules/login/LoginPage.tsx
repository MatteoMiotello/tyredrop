import React from "react";
import {Button, Form, Image, Input, Layout, Space} from "antd";
import {Content, Footer} from "antd/es/layout/layout";
import logo from "../../assets/logo-extended.jpeg";
import {useTranslation} from "react-i18next";


export const LoginPage: React.FC = () => {
    const { t, i18n } = useTranslation(  )

    return <Space className="m-0" direction="vertical" style={{width: '100%'}}>
        <Layout className="h-screen">
            <Content className="h-full m-24 bg-white flex flex-col justify-center items-center">
                <img
                    src={logo}
                    className="w-72"
                    alt="Tires in the world logo"
                ></img>
                <Form
                    name="login"
                    layout="vertical"
                    size="large"
                    className="text-center"
                >
                    <Form.Item
                        name="username"
                        className="mb-3 w-full my-auto my-2"
                        rules={[{required: true, message:"Username is required"}]}
                    >
                        <Input placeholder="Username"/>
                    </Form.Item>
                    <Form.Item
                        name="password"
                        rules={[{required: true, message:"Password is required"}]}
                        className="mb-3 w-full my-auto my-2"
                    >
                        <Input.Password placeholder="Password"/>
                    </Form.Item>
                    <Form.Item>
                        <Button
                            htmlType="submit"
                            type="primary"
                            className="mx-2"
                        >
                            { t( 'login.Submit' ) }
                        </Button>
                    </Form.Item>
                </Form>
            </Content>
            <Footer>Footer</Footer>
        </Layout>
    </Space>
}