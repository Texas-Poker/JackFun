using System;
using JackFun.UI;
using UnityEngine;

namespace JackFun
{
    public class Main : MonoBehaviour
    {
        private void Awake()
        {
            UIUtil.Init();
            UILoading.Open();
            HttpManager.Init();
            NetPitaya.Init();
            HttpManager.Entry(() =>
            {
                UILoading.Close();
                UIUtil.OpenUI<UILogin>();
            });
        }

        private void OnDestroy()
        {
            NetPitaya.Release();
        }
    }
}