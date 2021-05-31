using System;
using FairyGUI;
using UnityEngine;

namespace JackFun.UI
{
    public class UITips : UIPanel
    {
        public static void Open(string str)
        {
            Instance.gameObject.SetActive(true);
            Instance.ShowTips(str);
        }

        private static UITips Instance { get; set; }

        private GTextField _labelContent;
        private Transition _tween;

        private void Awake()
        {
            Instance = this;
            packageName = "FP_Common";
            componentName = "FC_Tips";
            sortingOrder = 80;
        }

        private void Start()
        {
            _labelContent = ui.GetChild("Label") as GTextField;
            _tween = ui.GetTransition("move_tween");
            gameObject.SetActive(false);
        }


        /// <summary>
        /// 显示提示
        /// </summary>
        /// <param name="content"></param>
        private void ShowTips(string content)
        {
            if (_tween.playing)
            {
                _tween.Stop();
                return;
            }

            _labelContent.text = content;
            ui.position = new Vector3(0, 0, 0);
            _tween.Play(() => { gameObject.SetActive(false); });
        }
    }
}