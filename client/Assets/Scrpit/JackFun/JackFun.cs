using System;
using UnityEngine;

namespace JackFun
{
    public class JackFun: MonoBehaviour
    {
        private void Start()
        {
            HttpManager.Init();
        }
    }
}