using System;
using FairyGUI;

namespace JackFun.UI
{
    public class UILogin : UIPanel
    {
        private GTextInput _inputAccount;
        private GTextInput _inputPassword;
        private GButton _btnLogin;

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

        private void InitRoot()
        {
            _inputAccount = (GTextInput) ui.GetChild("AccountInput").asTextField;
            _inputPassword = (GTextInput) ui.GetChild("PasswordInput").asTextField;
            _btnLogin = ui.GetChild("BtnLogin").asButton;
        }

        private void AddEvent()
        {
            _btnLogin.onClick.Add(ONBtnLoginClick);
        }

        // ReSharper disable Unity.PerformanceAnalysis
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

            HttpManager.Login(_inputAccount.text, _inputPassword.text);
        }
    }
}