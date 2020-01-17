import { Button, Col, Form, Input, Popover, Progress, Row, Select, message } from 'antd';
import { FormattedMessage, formatMessage } from 'umi-plugin-react/locale';
import React, { Component, useState, useRef, useEffect } from 'react';
import { Dispatch } from 'redux';
import { FormComponentProps } from 'antd/es/form';
import Link from 'umi/link';
import { connect } from 'dva';
import router from 'umi/router';

import { StateType } from './model';
import styles from './style.less';

const FormItem = Form.Item;
const { Option } = Select;
const InputGroup = Input.Group;

const passwordStatusMap = {
  ok: (
    <div className={styles.success}>
      <FormattedMessage id="useranduserregister.strength.strong" />
    </div>
  ),
  pass: (
    <div className={styles.warning}>
      <FormattedMessage id="useranduserregister.strength.medium" />
    </div>
  ),
  poor: (
    <div className={styles.error}>
      <FormattedMessage id="useranduserregister.strength.short" />
    </div>
  ),
};

const passwordProgressMap: {
  ok: 'success';
  pass: 'normal';
  poor: 'exception';
} = {
  ok: 'success',
  pass: 'normal',
  poor: 'exception',
};

interface UserRegisterProps extends FormComponentProps {
  dispatch: Dispatch<any>;
  userAndUserRegister: StateType;
  submitting: boolean;
}

export interface UserRegisterParams {
  mobile: string;
  password: string;
  confirm: string;
  captcha: string;
  prefix: string;
}

const UserRegister:React.FC<UserRegisterProps> = props => {
  const { userAndUserRegister, form ,dispatch, submitting} = props;
  const [time, setTime] = useState(0)
  const [confirmDirty, setConfirmDirty] = useState(false)
  const [visible, setVisible] = useState(false)
  const [help, setHelp] = useState('')
  const [prefix, setPrefix] = useState('86')

  const interval = useRef<number>();

  useEffect(()=> {
    const account = form.getFieldValue('mail');
    if (userAndUserRegister.status === 'ok') {
      message.success('注册成功！');
      router.push({
        pathname: '/user/register-result',
        state: {
          account,
        },
      });
    }
  }, [userAndUserRegister.status])

  useEffect(()=> () =>clearInterval(interval.current),[])


  const onGetCaptcha = () => {
    let currentTime = 59;
    setTime(currentTime)
    interval.current = window.setInterval(() => {
      currentTime -= 1;
      setTime(currentTime)
      if (currentTime === 0) {
        clearInterval(interval.current);
      }
    }, 1000);
  };

  const getPasswordStatus = () => {
    const value = form.getFieldValue('password');
    if (value && value.length > 9) {
      return 'ok';
    }
    if (value && value.length > 5) {
      return 'pass';
    }
    return 'poor';
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    form.validateFields({ force: true }, (err, values) => {
      if (!err) {
        dispatch({
          type: 'userAndUserRegister/submit',
          payload: {
            ...values,
            prefix,
          },
        });
      }
    });
  };

  const checkConfirm = (rule: any, value: string, callback: (messgae?: string) => void) => {
    if (value && value !== form.getFieldValue('password')) {
      callback(formatMessage({ id: 'useranduserregister.password.twice' }));
    } else {
      callback();
    }
  };

  const checkPassword = (rule: any, value: string, callback: (messgae?: string) => void) => {
    if (!value) {
      setHelp(formatMessage({ id: 'useranduserregister.password.required' }))
      setVisible(!!value)
      callback('error');
    } else {
      setHelp('')
      if (!visible) {
      setVisible(!!value)
      }
      if (value.length < 6) {
        callback('error');
      } else {
        if (value && confirmDirty) {
          form.validateFields(['confirm'], { force: true });
        }
        callback();
      }
    }
  };

  const changePrefix = (value: string) => {
    setPrefix(value)
  };

  const renderPasswordProgress = () => {
    const value = form.getFieldValue('password');
    const passwordStatus =getPasswordStatus();
    return value && value.length ? (
      <div className={styles[`progress-${passwordStatus}`]}>
        <Progress
          status={passwordProgressMap[passwordStatus]}
          className={styles.progress}
          strokeWidth={6}
          percent={value.length * 10 > 100 ? 100 : value.length * 10}
          showInfo={false}
        />
      </div>
    ) : null;
  };

    const { getFieldDecorator } = form;

    return (
      <div className={styles.main}>
        <h3>
          <FormattedMessage id="useranduserregister.register.register" />
        </h3>
        <Form onSubmit={handleSubmit}>
          {/* <FormItem>
            {getFieldDecorator('mail', {
              rules: [
                {
                  required: true,
                  message: formatMessage({ id: 'useranduserregister.email.required' }),
                },
                {
                  type: 'email',
                  message: formatMessage({ id: 'useranduserregister.email.wrong-format' }),
                },
              ],
            })(
              <Input
                size="large"
                placeholder={formatMessage({ id: 'useranduserregister.email.placeholder' })}
              />,
            )}
          </FormItem> */}
          <FormItem>
            <InputGroup compact>
              <Select
                size="large"
                value={prefix}
                onChange={changePrefix}
                style={{ width: '20%' }}
              >
                <Option value="86">+86</Option>
                <Option value="87">+87</Option>
              </Select>
              {getFieldDecorator('mobile', {
                rules: [
                  {
                    required: true,
                    message: formatMessage({ id: 'useranduserregister.phone-number.required' }),
                  },
                  {
                    pattern: /^\d{11}$/,
                    message: formatMessage({ id: 'useranduserregister.phone-number.wrong-format' }),
                  },
                ],
              })(
                <Input
                  size="large"
                  style={{ width: '80%' }}
                  placeholder={formatMessage({ id: 'useranduserregister.phone-number.placeholder' })}
                />,
              )}
            </InputGroup>
          </FormItem>
          <FormItem>
            <Row gutter={8}>
              <Col span={16}>
                {getFieldDecorator('captcha', {
                  rules: [
                    {
                      required: true,
                      message: formatMessage({ id: 'useranduserregister.verification-code.required' }),
                    },
                  ],
                })(
                  <Input
                    size="large"
                    placeholder={formatMessage({ id: 'useranduserregister.verification-code.placeholder' })}
                  />,
                )}
              </Col>
              <Col span={8}>
                <Button
                  size="large"
                  disabled={!!time}
                  className={styles.getCaptcha}
                  onClick={onGetCaptcha}
                >
                  {time
                    ? `${time} s`
                    : formatMessage({ id: 'useranduserregister.register.get-verification-code' })}
                </Button>
              </Col>
            </Row>
          </FormItem>
          <FormItem help={help}>
            <Popover
              getPopupContainer={node => {
                if (node && node.parentNode) {
                  return node.parentNode as HTMLElement;
                }
                return node;
              }}
              content={
                <div style={{ padding: '4px 0' }}>
                  {passwordStatusMap[getPasswordStatus()]}
                  {renderPasswordProgress()}
                  <div style={{ marginTop: 10 }}>
                    <FormattedMessage id="useranduserregister.strength.msg" />
                  </div>
                </div>
              }
              overlayStyle={{ width: 240 }}
              placement="right"
              visible={visible}
            >
              {getFieldDecorator('password', {
                rules: [
                  {
                    validator: checkPassword,
                  },
                ],
              })(
                <Input
                  size="large"
                  type="password"
                  placeholder={formatMessage({ id: 'useranduserregister.password.placeholder' })}
                />,
              )}
            </Popover>
          </FormItem>
          <FormItem>
            {getFieldDecorator('confirm', {
              rules: [
                {
                  required: true,
                  message: formatMessage({ id: 'useranduserregister.confirm-password.required' }),
                },
                {
                  validator: checkConfirm,
                },
              ],
            })(
              <Input
                size="large"
                type="password"
                placeholder={formatMessage({ id: 'useranduserregister.confirm-password.placeholder' })}
              />,
            )}
          </FormItem>
          <FormItem>
            <Button
              size="large"
              loading={submitting}
              className={styles.submit}
              type="primary"
              htmlType="submit"
            >
              <FormattedMessage id="useranduserregister.register.register" />
            </Button>
            <Link className={styles.login} to="/user/login">
              <FormattedMessage id="useranduserregister.register.sign-in" />
            </Link>
          </FormItem>
        </Form>
      </div>
    );
}

export default connect(
  ({
    userAndUserRegister,
    loading,
  }: {
    userAndUserRegister: StateType;
    loading: {
      effects: {
        [key: string]: boolean;
      };
    };
  }) => ({
    userAndUserRegister,
    submitting: loading.effects['userAndUserRegister/submit'],
  }),
)(Form.create<UserRegisterProps>()(UserRegister));
