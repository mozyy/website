import { UserInfo } from '../../../../proto/user/user_pb'
import { UserServiceClient } from '../../../../proto/user/UserServiceClientPb'

import request from '@/utils/request';
import { UserRegisterParams } from './index';

const client = new UserServiceClient('http://localhost:8080')

export const register = async(params: UserRegisterParams) => {
  const userInfo  = new UserInfo()
  userInfo.setMobile(params.mobile)
  userInfo.setPassword(params.password)
  client.register(userInfo,{},(err, response)=>{
    if (err) {
      console.log(222,'error')
    }else{
      console.log(1111,response.getMessage())
    }
  })
}


export async function fakeRegister(params: UserRegisterParams) {
  return request('/api/register', {
    method: 'POST',
    data: params,
  });
}
