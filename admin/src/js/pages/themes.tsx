import Layout from 'components/layout';
import {Table} from 'components/table';
import Theme from 'lib/theme';
import * as React from 'react';
import {Link} from 'react-router-dom';

interface State {
  themes: Theme[];
  loading: boolean;
}

export default class ThemesPage extends React.Component<{}, State> {
  constructor(props: any) {
    super(props);

    this.state = {
      themes: [],
      loading: true,
    };
  }

  componentDidMount() {
    Theme.list().then((themes) => {
      this.setState({
        loading: false,
        themes: themes,
      });
    });
  }

  render() {
    return (
      <Layout className="themes">
        <header>
          <Link className="button button--green button--center" to="/themes-install">
            Get More
          </Link>
          <h1>Themes</h1>
        </header>

        <h2>Installed themes</h2>
        <Table loading={this.state.loading}>
          {this.state.themes.map((theme) => {
            return (
              <Link key={theme.name} className="tr" to={`/themes/${theme.name}`}>
                {theme.name || 'untitled'}
              </Link>
            );
          })}
        </Table>
      </Layout>
    );
  }
}
