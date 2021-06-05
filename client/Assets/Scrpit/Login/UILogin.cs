using System;
using FairyGUI;
using JackFun.Lobby;

namespace JackFun.UI
{
    public class UILogin : UIPanel
    {
        private GTextInput _inputAccount;
        private GTextInput _inputPassword;
        private GButton _btnLogin;
        private GButton _btnToRegister;
        private Controller _controller;
        private GButton _btnToLogin;
        private GButton _btnRegister;

        private void Awake()
        {
            packageName = "FP_Login_Register";
            componentName = "FC_Login_Register";
        }

        private void Start()
        {
            InitRoot();
            AddEvent();
        }

        private void OnDestroy()
        {
            _btnLogin.onClick.Clear();
            _btnRegister.onClick.Clear();
            _btnToRegister.onClick.Clear();
            _btnToLogin.onClick.Clear();
        }

        private void InitRoot()
        {
            _inputAccount = (GTextInput) ui.GetChild("AccountInput").asTextField;
            _inputPassword = (GTextInput) ui.GetChild("PasswordInput").asTextField;
            _btnLogin = ui.GetChild("BtnLogin").asButton;
            _btnRegister = ui.GetChild("BtnRegister").asButton;
            _btnToRegister = ui.GetChild("BtnToRegister").asButton;
            _btnToLogin = ui.GetChild("BtnToLogin").asButton;
            _controller = ui.GetController("status");
        }

        private void AddEvent()
        {
            _btnLogin.onClick.Add(ONBtnLoginClick);
            _btnRegister.onClick.Add(ONBtnRegisterClick);
            _btnToRegister.onClick.Add(ONBtnToRegisterClick);
            _btnToLogin.onClick.Add(ONBtnToLoginClick);
        }


        private void ONBtnLoginClick()
        {
            if (String.IsNullOrEmpty(_inputAccount.text))
            {
                UITips.Open("账号不能为空！");
                return;
            }

            if (String.IsNullOrEmpty(_inputPassword.text))
            {
                UITips.Open("密码不能为空！");
                return;
            }

            HttpManager.Login(_inputAccount.text, _inputPassword.text,LoginController.CallAuth);
        }

        private void ONBtnToRegisterClick()
        {
            _controller.selectedPage = "register";
        }

        private void ONBtnToLoginClick()
        {
            _controller.selectedPage = "login";
        }

        private void ONBtnRegisterClick()
        {
            if (String.IsNullOrEmpty(_inputAccount.text))
            {
                UITips.Open("账号不能为空！");
                return;
            }

            if (String.IsNullOrEmpty(_inputPassword.text))
            {
                UITips.Open("密码不能为空！");
                return;
            }

            HttpManager.Register(_inputAccount.text, _inputPassword.text);
        }
    }
}