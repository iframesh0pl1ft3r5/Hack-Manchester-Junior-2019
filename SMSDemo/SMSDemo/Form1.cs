using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using System.Windows.Forms;
using Clockwork;
using MaterialSkin;
using MaterialSkin.Controls;


namespace SMSDemo
{
    public partial class Form1 : MaterialForm
    {
        public Form1()
        {
            InitializeComponent();
            //Initialize accent colours and themes
            var materialSkinManager = MaterialSkinManager.Instance;
            materialSkinManager.AddFormToManage(this);
            materialSkinManager.Theme = MaterialSkinManager.Themes.DARK;
            materialSkinManager.ColorScheme = new ColorScheme(Primary.Grey800, Primary.Grey900, Primary.Grey500, Accent.Cyan100, TextShade.WHITE);
        }

        private void Form1_Load(object sender, EventArgs e)
        {
            //Double buffer graphics (prevent flickering)
            DoubleBuffered = true;
            SetStyle(ControlStyles.OptimizedDoubleBuffer, true);

        }

        private void MaterialLabel2_Click(object sender, EventArgs e)
        {

        }
        protected override CreateParams CreateParams
        {
            //Double buffer graphics (prevent flickering)
            get
            {
                CreateParams handleParam = base.CreateParams;
                handleParam.ExStyle |= 0x02000000;   // WS_EX_COMPOSITED       
                return handleParam;
            }
        }
        private async void FadeIn(Form o, int interval = 80)
        {
            //Object is not fully invisible. Fade it in
            while (o.Opacity < 1.0)
            {
                await Task.Delay(interval);
                o.Opacity += 0.05;
            }
            o.Opacity = 1; //make fully visible       
        }

        private async void FadeOut(Form o, int interval = 80)
        {
            //Object is fully visible. Fade it out
            while (o.Opacity > 0.0)
            {
                await Task.Delay(interval);
                o.Opacity -= 0.05;
            }
            o.Opacity = 0; //make fully invisible       
        }
        private void MaterialFlatButton1_Click(object sender, EventArgs e)
        {
            materialFlatButton1.Visible = false;
            materialDivider4.Visible = true;
            materialDivider5.Visible = true;
            materialDivider6.Visible = true;
            label7.Visible = true;
            label8.Visible = true;
            label9.Visible = true;
            label10.Visible = true;
            label11.Visible = true;
            label12.Visible = true;
            materialRaisedButton6.Visible = true;
            materialRaisedButton7.Visible = true;
            materialRaisedButton8.Visible = true;
            materialRaisedButton9.Visible = true;
            materialRaisedButton10.Visible = true;
            materialRaisedButton11.Visible = true;
            Size tempSize = materialDivider7.Size;
            tempSize.Width = 158;
            tempSize.Height = 390;
            materialDivider7.Size = tempSize;


        }

        private void MaterialRaisedButton1_Click(object sender, EventArgs e)
        {
            if (MessageBox.Show("Are you sure you want to delete this class?", "Confirm", MessageBoxButtons.YesNo, MessageBoxIcon.Question) == DialogResult.Yes)
            {
                // user clicked yes

                materialDivider2.Location = materialDivider1.Location;
                label3.Location = label2.Location;
                label4.Location = label1.Location;
                materialRaisedButton2.Location = gotoHistory.Location;
                materialRaisedButton3.Location = removeHistory.Location;
                label1.Visible = false;
                label2.Visible = false;
                materialDivider1.Visible = false;
                gotoHistory.Visible = false;
                removeHistory.Visible = false;

            }
            else
            {
                // user clicked no

            }
        }

        private void GotoHistory_Click(object sender, EventArgs e)
        {
            FadeOut(this, 100);
            var form2 = new Form2();
            form2.Closed += (s, args) => this.Close();
            form2.Show();
            Console.Beep();


        }

    }
}
