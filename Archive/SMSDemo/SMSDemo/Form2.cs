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
    public partial class Form2 : MaterialForm
    {
        public Form2()
        {
            InitializeComponent();
            //Initialize accent colours and themes
            var materialSkinManager = MaterialSkinManager.Instance;
            materialSkinManager.AddFormToManage(this);
            materialSkinManager.Theme = MaterialSkinManager.Themes.DARK;
            materialSkinManager.ColorScheme = new ColorScheme(Primary.Grey800, Primary.Grey900, Primary.Grey500, Accent.Cyan100, TextShade.WHITE);
        }

        private void Form2_Load(object sender, EventArgs e)
        {
            // Add the pet to our listview
            listView1.View = View.Details;

            listView1.Columns.Add("Name");
            ListViewItem lvi1 = new ListViewItem();
            ListViewItem lvi2 = new ListViewItem();
            ListViewItem lvi3 = new ListViewItem();
            ListViewItem lvi4 = new ListViewItem();
            ListViewItem lvi5 = new ListViewItem();
            ListViewItem lvi6 = new ListViewItem();
            ListViewItem lvi7 = new ListViewItem();
            ListViewItem lvi8 = new ListViewItem();
            ListViewItem lvi9 = new ListViewItem();
            ListViewItem lvi10 = new ListViewItem();
            ListViewItem lvi11 = new ListViewItem();
            ListViewItem lvi12 = new ListViewItem();
            ListViewItem lvi13 = new ListViewItem();
            ListViewItem lvi14 = new ListViewItem();
            ListViewItem lvi15 = new ListViewItem();
            ListViewItem lvi16 = new ListViewItem();
            ListViewItem lvi17 = new ListViewItem();
            ListViewItem lvi18 = new ListViewItem();
            ListViewItem lvi19 = new ListViewItem();
            ListViewItem lvi20 = new ListViewItem();
            ListViewItem lvi21 = new ListViewItem();
            ListViewItem lvi22 = new ListViewItem();
            ListViewItem lvi23 = new ListViewItem();
            ListViewItem lvi24 = new ListViewItem();
            ListViewItem lvi25 = new ListViewItem();

            lvi1.Text = "Matthew O";
            lvi2.Text = "Jake R";
            lvi3.Text = "Joshua G";
            lvi4.Text = "Abraham W";
            lvi5.Text = "Kien A";
            lvi6.Text = "Karan H";
            lvi7.Text = "Blake N";
            lvi8.Text = "Oscar F";
            lvi9.Text = "Fred R";
            lvi10.Text = "John S";
            lvi11.Text = "Alan L";
            lvi12.Text = "Felix K";
            lvi13.Text = "Logan P";
            lvi14.Text = "Jake P";
            lvi15.Text = "Josh M";
            lvi16.Text = "Ruby G";
            lvi17.Text = "John P";
            lvi18.Text = "Alex G";
            lvi19.Text = "Kalpana R";
            lvi20.Text = "Alfie S";
            lvi21.Text = "Emma P";
            lvi22.Text = "Mohammed Y";
            lvi23.Text = "Euan G";
            lvi24.Text = "Asal H";
            lvi25.Text = "Hamza K";



            listView1.Items.Add(lvi1);
            listView1.Items.Add(lvi2);
            listView1.Items.Add(lvi3);
            listView1.Items.Add(lvi4);
            listView1.Items.Add(lvi5);
            listView1.Items.Add(lvi6);
            listView1.Items.Add(lvi7);
            listView1.Items.Add(lvi8);
            listView1.Items.Add(lvi9);
            listView1.Items.Add(lvi10);
            listView1.Items.Add(lvi11);
            listView1.Items.Add(lvi12);
            listView1.Items.Add(lvi13);
            listView1.Items.Add(lvi14);
            listView1.Items.Add(lvi15);
            listView1.Items.Add(lvi16);
            listView1.Items.Add(lvi17);
            listView1.Items.Add(lvi18);
            listView1.Items.Add(lvi19);
            listView1.Items.Add(lvi20);
            listView1.Items.Add(lvi21);
            listView1.Items.Add(lvi22);
            listView1.Items.Add(lvi23);
            listView1.Items.Add(lvi24);
            listView1.Items.Add(lvi25);







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
    }

}
