
package definitions

func init(){
  add(`AccountDetails`, &defAccountDetails{})
}

type defAccountDetails struct{}

func (*defAccountDetails) String() string {
	return `
<interface>
  <object class="GtkDialog" id="AccountDetailsDialog">
    <property name="title" translatable="yes">Account Details</property>
    <signal name="close" handler="on_close_signal" />
    <child internal-child="vbox">
      <object class="GtkBox" id="Vbox">
        <property name="homogeneous">false</property>
        <property name="orientation">GTK_ORIENTATION_VERTICAL</property>
        <child>
          <object class="GtkLabel" id="AccountMessageLabel">
            <property name="label" translatable="yes">Your account (for example: kim42@dukgo.com)</property>
          </object>
          <packing>
            <property name="expand">false</property>
            <property name="fill">true</property>
            <property name="position">0</property>
          </packing>
        </child>
        <child>
          <object class="GtkEntry" id="account">
            <property name="has-focus">true</property>
            <signal name="activate" handler="on_save_signal" />
          </object>
          <packing>
            <property name="expand">false</property>
            <property name="fill">true</property>
            <property name="position">1</property>
          </packing>
        </child>
        <child>
          <object class="GtkLabel" id="PasswordLabel">
            <property name="label" translatable="yes">Password</property>
          </object>
          <packing>
            <property name="expand">false</property>
            <property name="fill">true</property>
            <property name="position">2</property>
          </packing>
        </child>
        <child>
          <object class="GtkEntry" id="password">
            <property name="visibility">false</property>
            <signal name="activate" handler="on_save_signal" />
          </object>
          <packing>
            <property name="expand">false</property>
            <property name="fill">true</property>
            <property name="position">3</property>
          </packing>
        </child>
        <child>
          <object class="GtkButton" id="save">
            <property name="label" translatable="yes">Save</property>
            <signal name="clicked" handler="on_save_signal"/>
          </object>
          <packing>
            <property name="expand">false</property>
            <property name="fill">true</property>
            <property name="position">4</property>
          </packing>
        </child>
      </object>
    </child>
  </object>
</interface>

`
}