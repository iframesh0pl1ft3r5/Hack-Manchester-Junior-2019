Public Class Form1
    Private Sub Form1_Load(sender As Object, e As EventArgs) Handles MyBase.Load
        ResizeWindow()
        add_group("test")
    End Sub

    Sub ResizeWindow()
        ListView1.Top = 64
        ListView1.Left = 0
        ListView1.Height = Me.Height - 65
        ListView1.Width = Me.Width - 1
        MaterialRaisedButton1.Left = Me.Width - 167
        MaterialRaisedButton1.Top = 25

    End Sub

    Private Sub Form1_ResizeEnd(sender As Object, e As EventArgs) Handles Me.ResizeEnd
        ResizeWindow()
    End Sub

    Sub add_group(ByVal groupname As String)
        ListView1.Items.Add(groupname, ImageList1.Images.Count - 2)
    End Sub
End Class
