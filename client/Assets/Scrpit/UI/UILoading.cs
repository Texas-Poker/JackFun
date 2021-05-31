using System;
using FairyGUI;
using UnityEngine;

namespace JackFun.UI
{
    public class UILoading : UIPanel
    {
        public static void Open()
        {
            Instance.gameObject.SetActive(true);
        }

        public static void Close()
        {
            Instance.gameObject.SetActive(false);
        }

        private static UILoading Instance { get; set; }

        private void Awake()
        {
            Instance = this;
            packageName = "FP_Common";
            componentName = "FC_Loading";
            sortingOrder = 100;
        }
    }
}